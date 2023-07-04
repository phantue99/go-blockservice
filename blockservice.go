// package blockservice implements a BlockService interface that provides
// a single GetBlock/AddBlock interface that seamlessly retrieves data either
// locally or from a remote peer through the exchange.
package blockservice

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-blockservice/tikv"
	cid "github.com/ipfs/go-cid"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	exchange "github.com/ipfs/go-ipfs-exchange-interface"
	ipld "github.com/ipfs/go-ipld-format"
	logging "github.com/ipfs/go-log/v2"
	"github.com/ipfs/go-verifcid"

	"github.com/ipfs/go-blockservice/internal"
)

var logger = logging.Logger("blockservice")

// BlockGetter is the common interface shared between blockservice sessions and
// the blockservice.
type BlockGetter interface {
	// GetBlock gets the requested block.
	GetBlock(ctx context.Context, c cid.Cid) (blocks.Block, error)

	// GetBlocks does a batch request for the given cids, returning blocks as
	// they are found, in no particular order.
	//
	// It may not be able to find all requested blocks (or the context may
	// be canceled). In that case, it will close the channel early. It is up
	// to the consumer to detect this situation and keep track which blocks
	// it has received and which it hasn't.
	GetBlocks(ctx context.Context, ks []cid.Cid) <-chan blocks.Block
}

// BlockService is a hybrid block datastore. It stores data in a local
// datastore and may retrieve data from a remote Exchange.
// It uses an internal `datastore.Datastore` instance to store values.
type BlockService interface {
	io.Closer
	BlockGetter

	// Blockstore returns a reference to the underlying blockstore
	Blockstore() blockstore.Blockstore

	// Exchange returns a reference to the underlying exchange (usually bitswap)
	Exchange() exchange.Interface

	// AddBlock puts a given block to the underlying datastore
	AddBlock(ctx context.Context, o blocks.Block) error

	// AddBlocks adds a slice of blocks at the same time using batching
	// capabilities of the underlying datastore whenever possible.
	AddBlocks(ctx context.Context, bs []blocks.Block) error

	// DeleteBlock deletes the given block from the blockservice.
	DeleteBlock(ctx context.Context, o cid.Cid) error
}

type blockService struct {
	blockstore blockstore.Blockstore
	exchange   exchange.Interface
	// If checkFirst is true then first check that a block doesn't
	// already exist to avoid republishing the block on the exchange.
	checkFirst bool
}

type fileRecord struct {
	FileRecordID string
	Size         uint64
}

type fileInfo struct {
	FileRecordID string
	Size         uint64
	Offset       uint64
}

var (
	uploader           string
	pinningService     string
	apiKey             string
	isDedicatedGateway bool
)

func InitBlockService(uploaderURL, pinningServiceURL, _apiKey string, _isDedicatedGateway bool) error {
	if uploaderURL != "" {
		uploader = uploaderURL
	}
	if pinningServiceURL != "" {
		pinningService = pinningServiceURL
	}
	if _apiKey != "" {
		apiKey = _apiKey
	}
	isDedicatedGateway = _isDedicatedGateway

	// Return an error if any of the URLs is empty.
	if uploader == "" || pinningService == "" || apiKey == "" {
		return errors.New("error: empty url or api key")
	}

	return nil
}

// NewBlockService creates a BlockService with given datastore instance.
func New(bs blockstore.Blockstore, rem exchange.Interface) BlockService {
	if rem == nil {
		logger.Debug("blockservice running in local (offline) mode.")
	}

	return &blockService{
		blockstore: bs,
		exchange:   rem,
		checkFirst: true,
	}
}

// NewWriteThrough creates a BlockService that guarantees writes will go
// through to the blockstore and are not skipped by cache checks.
func NewWriteThrough(bs blockstore.Blockstore, rem exchange.Interface) BlockService {
	if rem == nil {
		logger.Debug("blockservice running in local (offline) mode.")
	}

	return &blockService{
		blockstore: bs,
		exchange:   rem,
		checkFirst: false,
	}
}

// Blockstore returns the blockstore behind this blockservice.
func (s *blockService) Blockstore() blockstore.Blockstore {
	return s.blockstore
}

// Exchange returns the exchange behind this blockservice.
func (s *blockService) Exchange() exchange.Interface {
	return s.exchange
}

// NewSession creates a new session that allows for
// controlled exchange of wantlists to decrease the bandwidth overhead.
// If the current exchange is a SessionExchange, a new exchange
// session will be created. Otherwise, the current exchange will be used
// directly.
func NewSession(ctx context.Context, bs BlockService) *Session {
	exch := bs.Exchange()
	if sessEx, ok := exch.(exchange.SessionExchange); ok {
		return &Session{
			sessCtx:  ctx,
			ses:      nil,
			sessEx:   sessEx,
			bs:       bs.Blockstore(),
			notifier: exch,
		}
	}
	return &Session{
		ses:      exch,
		sessCtx:  ctx,
		bs:       bs.Blockstore(),
		notifier: exch,
	}
}

type TikvUser struct {
	FileRecordID string
}

type ZipReader struct {
	File []File
}
type File struct {
	Name               string `json:"Name"`
	CompressedSize     uint32 `json:"CompressedSize"`
	UncompressedSize   uint32 `json:"UncompressedSize"`
	CompressedSize64   uint64 `json:"CompressedSize64"`
	UncompressedSize64 uint64 `json:"UncompressedSize64"`
	Offset             uint64 `json:"Offset"`
}

// AddBlock adds a particular block to the service, Putting it into the datastore.
func (s *blockService) AddBlock(ctx context.Context, o blocks.Block) error {
	ctx, span := internal.StartSpan(ctx, "blockService.AddBlock")
	defer span.End()

	err := AddBlock(ctx, o, s.checkFirst)
	if err != nil {
		return err
	}

	if s.exchange != nil {
		if err := s.exchange.NotifyNewBlocks(ctx, o); err != nil {
			logger.Errorf("NotifyNewBlocks: %s", err.Error())
		}
	}

	return nil
}

func AddBlock(ctx context.Context, o blocks.Block, checkFirst bool) error {
	var fr fileRecord
	userID, _ := ctx.Value("userID").(string)
	if userID != "" {
		userKV, err := tikv.Get([]byte(userID))
		if err != nil {
			return nil
		} else {
			if err := json.Unmarshal(userKV.V, &fr); err != nil {
				return err
			}
		}
	}

	c := o.Cid()
	// hash security
	if err := verifcid.ValidateCid(c); err != nil {
		return err
	}

	if checkFirst {
		_, err := tikv.Get(c.Hash())
		if err == nil {
			return nil
		}
	}

	hash, err := internal.GetHashStringFromCid(c.String())
	if err != nil {
		return err
	}
	// Create a temporary file with the name being the hash of the CID
	tmpFile, err := os.Create(hash)
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer os.Remove(tmpFile.Name()) // clean up the temporary file

	// Write the block data to the temporary file
	_, err = tmpFile.Write(o.RawData())
	if err != nil {
		return fmt.Errorf("failed to write block data to temporary file: %w", err)
	}
	if err = tmpFile.Close(); err != nil {
		return fmt.Errorf("failed to close temporary file: %w", err)
	}

	var (
		fileRecordID = fr.FileRecordID
		lastSize     uint64
		files        []File
	)
	if fr.FileRecordID == "" || fr.Size > 100*1024*1024 {
		fileRecordID, files, lastSize, err = uploadFiles([]string{tmpFile.Name()})
		if err != nil {
			return fmt.Errorf("failed to upload file and get file record ID: %w", err)
		}
	} else {
		files, lastSize, err = appendFiles([]string{tmpFile.Name()}, fileRecordID)
		if err != nil {
			return fmt.Errorf("failed to upload file : %w", err)
		}
	}

	f := fileRecord{fileRecordID, lastSize}
	bf, err := json.Marshal(f)
	if err != nil {
		return fmt.Errorf("failed to marshal `fileInfo`: %w", err)
	}
	if err := tikv.Puts([]byte(userID), []byte(bf)); err != nil {
		return fmt.Errorf("failed to put data in TiKV: %w", err)
	}

	for _, f := range files {
		if strings.Contains(f.Name, o.Cid().Hash().String()) {
			fInfo := fileInfo{fileRecordID, f.CompressedSize64, f.Offset}
			fInfoBytes, err := json.Marshal(fInfo)
			if err != nil {
				return err
			}
			if err := tikv.Puts(o.Cid().Hash(), []byte(fInfoBytes)); err != nil {
				return err
			}
			break
		}
	}

	logger.Debugf("BlockService.BlockAdded %s", c)

	return nil
}

func uploadFiles(files []string) (string, []File, uint64, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	// Iterate over files and add them as form parts
	for _, file := range files {
		part, err := writer.CreateFormFile("file", filepath.Base(file))
		if err != nil {
			return "", nil, 0, fmt.Errorf("failed to create form file: %w", err)
		}
		f, err := os.Open(file)
		if err != nil {
			return "", nil, 0, fmt.Errorf("failed to open file %s: %w", file, err)
		}
		defer f.Close()
		if _, err = io.Copy(part, f); err != nil {
			return "", nil, 0, fmt.Errorf("failed to copy file content: %w", err)
		}
	}

	// Close the multipart form writer
	if err := writer.Close(); err != nil {
		return "", nil, 0, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/packUpload", uploader), body)
	if err != nil {
		return "", nil, 0, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", nil, 0, fmt.Errorf("failed to post raw data: %w", err)
	}
	defer resp.Body.Close()
	type FileRecord struct {
		ID       string `json:"ID"`
		Owner    string `json:"owner"`
		Name     string `json:"name"`
		Size     int    `json:"size"`
		ReaderID int    `json:"readerId"`
	}
	type Response struct {
		FileRecord FileRecord
		ZipReader  ZipReader
	}
	var (
		response     Response
		fileRecordID string
		size         uint64
	)
	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return "", nil, 0, fmt.Errorf("failed to decode response body: %w", err)
		}
		fileRecordID = response.FileRecord.ID
		size = response.ZipReader.File[len(files)-1].Offset + response.ZipReader.File[len(files)-1].UncompressedSize64
	}
	return fileRecordID, response.ZipReader.File, size, nil
}
func appendFiles(files []string, fileRecordId string) ([]File, uint64, error) {
	// Create new multipart form writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Iterate over files and add them as form parts
	for _, file := range files {
		part, err := writer.CreateFormFile("file", filepath.Base(file))
		if err != nil {
			return nil, 0, fmt.Errorf("failed to create form file: %w", err)
		}
		f, err := os.Open(file)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to open file %s: %w", file, err)
		}
		defer f.Close()
		if _, err = io.Copy(part, f); err != nil {
			return nil, 0, fmt.Errorf("failed to copy file content: %w", err)
		}
	}

	// Close the multipart form writer
	if err := writer.Close(); err != nil {
		return nil, 0, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	// Create new HTTP request and set headers
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/zipAction", uploader), body)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.URL.RawQuery = url.Values{
		"file_record_id": {fileRecordId},
		"action_type":    {"1"},
	}.Encode()
	req.Header.Set("Content-Type", writer.FormDataContentType())
	// Send request and handle response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, 0, fmt.Errorf("server returned status %d, reqURI: %s", resp.StatusCode, req.URL.String())
	}

	var (
		response      ZipReader
		lastFileIndex int
		lastSize      uint64
	)
	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return nil, 0, fmt.Errorf("failed to decode response body: %w", err)
		}
		lastFileIndex = len(response.File) - 1

		lastSize = response.File[lastFileIndex].Offset + response.File[lastFileIndex].UncompressedSize64

	}

	return response.File, lastSize, nil
}

type TikvBlock struct {
	FileName     []byte
	FileRecordID string
	Range        string
}

var lock = sync.RWMutex{}

func (s *blockService) AddBlocks(ctx context.Context, bs []blocks.Block) error {
	lock.Lock()
	defer lock.Unlock()
	ctx, span := internal.StartSpan(ctx, "blockService.AddBlocks")
	defer span.End()

	toput, err := AddBlocks(ctx, bs, s.checkFirst)
	if err != nil {
		return err
	}

	if s.exchange != nil {
		logger.Debugf("BlockService.BlockAdded %d blocks", len(toput))
		if err := s.exchange.NotifyNewBlocks(ctx, toput...); err != nil {
			logger.Errorf("NotifyNewBlocks: %s", err.Error())
		}
	}
	return nil
}

func AddBlocks(ctx context.Context, bs []blocks.Block, checkFirst bool) ([]blocks.Block, error) {
	var fr fileRecord
	var toput []blocks.Block

	userID, _ := ctx.Value("userID").(string)
	if userID != "" {
		userKV, err := tikv.Get([]byte(userID))
		if err != nil {
			return toput, nil
		} else {
			if err := json.Unmarshal(userKV.V, &fr); err != nil {
				return nil, err
			}
		}
	}

	// hash security
	for _, b := range bs {
		err := verifcid.ValidateCid(b.Cid())
		if err != nil {
			return nil, err
		}
	}
	if checkFirst {
		toput = make([]blocks.Block, 0, len(bs))
		for _, b := range bs {
			_, err := tikv.Get(b.Cid().Hash())
			if err == nil {
				continue // Skip already added block
			} else {
				toput = append(toput, b)
			}
		}
	} else {
		toput = bs
	}

	if len(toput) == 0 {
		return toput, nil
	}

	tempFiles := make([]string, len(toput))
	for i, b := range toput {
		tempFile, err := os.Create(b.Cid().Hash().String())
		if err != nil {
			return nil, err
		}
		defer os.Remove(tempFile.Name())
		_, err = tempFile.Write(b.RawData())
		if err != nil {
			return nil, err
		}
		tempFiles[i] = tempFile.Name()
	}

	var (
		fileRecordID = fr.FileRecordID
		lastSize     uint64
		err          error
		files        []File
	)
	if fr.FileRecordID == "" || fr.Size > 100*1024*1024 {
		fileRecordID, files, lastSize, err = uploadFiles(tempFiles)
		if err != nil {
			return nil, fmt.Errorf("failed to upload file and get file record ID: %w", err)
		}
	} else {
		files, lastSize, err = appendFiles(tempFiles, fileRecordID)
		if err != nil {
			return nil, err
		}
	}

	if userID != "" {
		f := fileRecord{fileRecordID, lastSize}
		bf, err := json.Marshal(f)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal `fileInfo`: %w", err)
		}
		if err := tikv.Puts([]byte(userID), []byte(bf)); err != nil {
			return nil, fmt.Errorf("failed to put data in TiKV: %w", err)
		}
	}

	for _, f := range files {
		for _, b := range toput {
			if strings.Contains(f.Name, b.Cid().Hash().String()) {
				fInfo := fileInfo{fileRecordID, f.CompressedSize64, f.Offset}
				fInfoBytes, err := json.Marshal(fInfo)
				if err != nil {
					return nil, err
				}
				if err := tikv.Puts(b.Cid().Hash(), []byte(fInfoBytes)); err != nil {
					return nil, err
				}
			}
		}
	}
	return toput, nil
}

// GetBlock retrieves a particular block from the service,
// Getting it from the datastore using the key (hash).
func (s *blockService) GetBlock(ctx context.Context, c cid.Cid) (blocks.Block, error) {
	ctx, span := internal.StartSpan(ctx, "blockService.GetBlock", trace.WithAttributes(attribute.Stringer("CID", c)))
	defer span.End()

	var f func() notifiableFetcher
	if s.exchange != nil {
		f = s.getExchange
	}

	return getBlock(ctx, c, s.blockstore, f) // hash security
}

func (s *blockService) GetUploader() (string, error) {
	if uploader != "" {
		return uploader, nil
	}
	return "", nil
}

func (s *blockService) getExchange() notifiableFetcher {
	return s.exchange
}

func getBlock(ctx context.Context, c cid.Cid, bs blockstore.Blockstore, fget func() notifiableFetcher) (blocks.Block, error) {
	err := verifcid.ValidateCid(c) // hash security
	if err != nil {
		return nil, err
	}

	kv1, err := tikv.Get(c.Hash())

	if err == nil {
		var f fileInfo

		if err := json.Unmarshal(kv1.V, &f); err != nil {
			return nil, err
		}

		endpoint, err := url.Parse(fmt.Sprintf("%s/cacheFile/%s", uploader, f.FileRecordID))
		if err != nil {
			return nil, err
		}

		rawQuery := endpoint.Query()
		rawQuery.Set("range", fmt.Sprintf("%d,%d", f.Offset, f.Size))
		endpoint.RawQuery = rawQuery.Encode()
		fileUrl := endpoint.String()

		fmt.Println(fileUrl)

		resp, err := http.Get(fileUrl)
		if err != nil {
			logger.Debugf("Failed to get data %v", err)
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			logger.Debugf("Request failed with status %d", resp.StatusCode)
			return nil, err
		}

		bdata, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Debugf("Failed read body %v", err)
			return nil, err
		}
		hash, err := internal.GetHashStringFromCid(c.String())
		if err != nil {
			logger.Debugf("GetHashFromCidString Error %v", err)
		}
		if isDedicatedGateway {
			addBandwidthUsage(f.Size, hash)
		}

		return blocks.NewBlockWithCid(bdata, c)
	}

	if ipld.IsNotFound(err) && fget != nil || err != nil {
		f := fget() // Don't load the exchange until we have to

		// TODO be careful checking ErrNotFound. If the underlying
		// implementation changes, this will break.
		logger.Debug("Blockservice: Searching bitswap")
		blk, err := f.GetBlock(ctx, c)
		if err != nil {
			return nil, err
		}
		cache := ctx.Value("cache")
		if cache != nil && cache == true {
			err = bs.Put(ctx, blk)
			if err != nil {
				return nil, err
			}
			err = AddBlock(ctx, blk, false)
			if err != nil {
				return nil, err
			}
			err = f.NotifyNewBlocks(ctx, blk)
			if err != nil {
				return nil, err
			}
		}
		logger.Debugf("BlockService.BlockFetched %s", c)
		return blk, nil
	}

	logger.Debug("Blockservice GetBlock: Not found")
	return nil, err
}

// GetBlocks gets a list of blocks asynchronously and returns through
// the returned channel.
// NB: No guarantees are made about order.
func (s *blockService) GetBlocks(ctx context.Context, ks []cid.Cid) <-chan blocks.Block {
	ctx, span := internal.StartSpan(ctx, "blockService.GetBlocks")
	defer span.End()

	var f func() notifiableFetcher
	if s.exchange != nil {
		f = s.getExchange
	}

	return getBlocks(ctx, ks, s.blockstore, f) // hash security
}

func getBlocks(ctx context.Context, ks []cid.Cid, bs blockstore.Blockstore, fget func() notifiableFetcher) <-chan blocks.Block {
	out := make(chan blocks.Block)

	go func() {
		defer close(out)

		allValid := true
		for _, c := range ks {
			if err := verifcid.ValidateCid(c); err != nil {
				allValid = false
				break
			}
		}

		if !allValid {
			ks2 := make([]cid.Cid, 0, len(ks))
			for _, c := range ks {
				// hash security
				if err := verifcid.ValidateCid(c); err == nil {
					ks2 = append(ks2, c)
				} else {
					logger.Errorf("unsafe CID (%s) passed to blockService.GetBlocks: %s", c, err)
				}
			}
			ks = ks2
		}

		var misses []cid.Cid
		for _, c := range ks {
			hit, err := getBlockCdn(ctx, c)
			if err != nil {
				misses = append(misses, c)
				continue
			}
			select {
			case out <- hit:
			case <-ctx.Done():
				return
			}
		}

		if len(misses) == 0 || fget == nil {
			return
		}

		f := fget() // don't load exchange unless we have to
		rblocks, err := f.GetBlocks(ctx, misses)
		if err != nil {
			logger.Debugf("Error with GetBlocks: %s", err)
			return
		}

		// batch available blocks together
		const batchSize = 32
		batch := make([]blocks.Block, 0, batchSize)
		for {
			var noMoreBlocks bool
		batchLoop:
			for len(batch) < batchSize {
				select {
				case b, ok := <-rblocks:
					if !ok {
						noMoreBlocks = true
						break batchLoop
					}

					logger.Debugf("BlockService.BlockFetched %s", b.Cid())
					batch = append(batch, b)
				case <-ctx.Done():
					return
				default:
					break batchLoop
				}
			}

			cache := ctx.Value("cache")

			if cache != nil && cache == true {
				// also write in the blockstore for caching, inform the exchange that the blocks are available
				err = bs.PutMany(ctx, batch)
				if err != nil {
					logger.Errorf("could not write blocks from the network to the blockstore: %s", err)
					return
				}
				_, err = AddBlocks(ctx, batch, false)
				if err != nil {
					logger.Errorf("could not add blocks from the network to the cdn: %s", err)
					return
				}

				err = f.NotifyNewBlocks(ctx, batch...)
				if err != nil {
					logger.Errorf("could not tell the exchange about new blocks: %s", err)
					return
				}
			}

			for _, b := range batch {
				select {
				case out <- b:
				case <-ctx.Done():
					return
				}
			}
			batch = batch[:0]
			if noMoreBlocks {
				break
			}
		}
	}()
	return out
}

// DeleteBlock deletes a block in the blockservice from the datastore
func (s *blockService) DeleteBlock(ctx context.Context, c cid.Cid) error {
	ctx, span := internal.StartSpan(ctx, "blockService.DeleteBlock", trace.WithAttributes(attribute.Stringer("CID", c)))
	defer span.End()

	err := s.blockstore.DeleteBlock(ctx, c)
	if err == nil {
		logger.Debugf("BlockService.BlockDeleted %s", c)
	}
	return err
}

func (s *blockService) Close() error {
	logger.Debug("blockservice is shutting down...")
	return s.exchange.Close()
}

type notifier interface {
	NotifyNewBlocks(context.Context, ...blocks.Block) error
}

// Session is a helper type to provide higher level access to bitswap sessions
type Session struct {
	bs       blockstore.Blockstore
	ses      exchange.Fetcher
	sessEx   exchange.SessionExchange
	sessCtx  context.Context
	notifier notifier
	lk       sync.Mutex
}

type notifiableFetcher interface {
	exchange.Fetcher
	notifier
}

type notifiableFetcherWrapper struct {
	exchange.Fetcher
	notifier
}

func (s *Session) getSession() notifiableFetcher {
	s.lk.Lock()
	defer s.lk.Unlock()
	if s.ses == nil {
		s.ses = s.sessEx.NewSession(s.sessCtx)
	}

	return notifiableFetcherWrapper{s.ses, s.notifier}
}

func (s *Session) getExchange() notifiableFetcher {
	return notifiableFetcherWrapper{s.ses, s.notifier}
}

func (s *Session) getFetcherFactory() func() notifiableFetcher {
	if s.sessEx != nil {
		return s.getSession
	}
	if s.ses != nil {
		// Our exchange isn't session compatible, let's fallback to non sessions fetches
		return s.getExchange
	}
	return nil
}
func getBlockCdn(ctx context.Context, c cid.Cid) (blocks.Block, error) {
	kv1, err := tikv.Get(c.Hash())
	if err == nil {
		var f fileInfo

		err := json.Unmarshal(kv1.V, &f)
		if err != nil {
			return nil, err
		}

		endpoint, err := url.Parse(fmt.Sprintf("%s/cacheFile/%s", uploader, f.FileRecordID))
		if err != nil {
			return nil, err
		}

		rawQuery := endpoint.Query()
		rawQuery.Set("range", fmt.Sprintf("%d,%d", f.Offset, f.Size))
		endpoint.RawQuery = rawQuery.Encode()
		fileUrl := endpoint.String()

		resp, err := http.Get(fileUrl)
		if err != nil {
			return nil, err
		}
		hash, err := internal.GetHashStringFromCid(c.String())
		if err != nil {
			logger.Debugf("GetHashFromCidString Error %v", err)
		}
		if isDedicatedGateway {
			addBandwidthUsage(f.Size, hash)
		}

		bdata, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			return blocks.NewBlockWithCid(bdata, c)
		}
	}
	return nil, err
}
func addBandwidthUsage(fileSize uint64, hash string) error {
	apiUrl := fmt.Sprintf("%s/api/hourlyUsage/bandwidth/", pinningService)
	reqBody, _ := json.Marshal(map[string]interface{}{
		"amount": fileSize,
		"cid":    hash,
	})
	client := &http.Client{}
	req, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer(reqBody))
	req.Header.Set("blockservice-API-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")
	_, err := client.Do(req)
	if err != nil {
		logger.Debugf("Failed to send Bandwidth Usage Error %v", err)
		return err
	}
	return nil
}

// GetBlock gets a block in the context of a request session
func (s *Session) GetBlock(ctx context.Context, c cid.Cid) (blocks.Block, error) {
	ctx, span := internal.StartSpan(ctx, "Session.GetBlock", trace.WithAttributes(attribute.Stringer("CID", c)))
	defer span.End()

	return getBlock(ctx, c, s.bs, s.getFetcherFactory()) // hash security
}

// GetBlocks gets blocks in the context of a request session
func (s *Session) GetBlocks(ctx context.Context, ks []cid.Cid) <-chan blocks.Block {
	ctx, span := internal.StartSpan(ctx, "Session.GetBlocks")
	defer span.End()

	return getBlocks(ctx, ks, s.bs, s.getFetcherFactory()) // hash security
}

var _ BlockGetter = (*Session)(nil)
