// package blockservice implements a BlockService interface that provides
// a single GetBlock/AddBlock interface that seamlessly retrieves data either
// locally or from a remote peer through the exchange.
package blockservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
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
var uploader string

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

type fileInfo struct {
	FileId string
	Size   int
}

func InitUploader(ip string) {
	if ip != "" {
		uploader = ip
	}
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

// AddBlock adds a particular block to the service, Putting it into the datastore.
func (s *blockService) AddBlock(ctx context.Context, o blocks.Block) error {
	ctx, span := internal.StartSpan(ctx, "blockService.AddBlock")
	defer span.End()

	c := o.Cid()
	// hash security
	err := verifcid.ValidateCid(c)
	if err != nil {
		return err
	}

	if s.checkFirst {
		_, err = tikv.Get(c.Bytes())
		if err == nil {
			return nil
		}
	}

	hash, err := internal.GetHashFromCidString(c.String())
	if err != nil {
		return err
	}

    resp, err := uploadRequest(hash, o.RawData())
    if err != nil {
        return fmt.Errorf("failed to post raw data: %w", err)
    }
    defer resp.Body.Close()

	var value string
	if resp.StatusCode == http.StatusOK {
		var res map[string]interface{}
        if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
            return fmt.Errorf("failed to decode response body: %w", err)
        }
		value = res["fileId"].(string)
	}

	f := fileInfo{value, len(o.RawData())}
	bf, err := json.Marshal(f)
    if err != nil {
        return fmt.Errorf("failed to marshal `fileInfo`: %w", err)
    }
	// set
    if err := tikv.Puts(c.Bytes(), []byte(bf)); err != nil {
        return fmt.Errorf("failed to put data in TiKV: %w", err)
    }

	logger.Debugf("BlockService.BlockAdded %s", c)

	if s.exchange != nil {
		if err := s.exchange.NotifyNewBlocks(ctx, o); err != nil {
			logger.Errorf("NotifyNewBlocks: %s", err.Error())
		}
	}

	return nil
}

func (s *blockService) AddBlocks(ctx context.Context, bs []blocks.Block) error {
	ctx, span := internal.StartSpan(ctx, "blockService.AddBlocks")
	defer span.End()

	// hash security
	for _, b := range bs {
		err := verifcid.ValidateCid(b.Cid())
		if err != nil {
			return err
		}
	}
	var toput []blocks.Block
	if s.checkFirst {
		toput = make([]blocks.Block, 0, len(bs))
		for _, b := range bs {
			_, err := tikv.Get(b.Cid().Bytes())
			if err == nil {
				return nil
			} else {
				toput = append(toput, b)
			}
		}
	} else {
		toput = bs
	}

	if len(toput) == 0 {
		return nil
	}

	var wg sync.WaitGroup
    var mu sync.Mutex // Protect err to avoid data race.
	errc := make(chan error, len(toput))

	for _, b := range toput {
		wg.Add(1)

		go func(b blocks.Block) {
			defer wg.Done()
			_, err := tikv.Get(b.Cid().Bytes())
			if err == nil {
				return
			}

			hash, err := internal.GetHashFromCidString(b.String())
			if err != nil {
                mu.Lock()
                errc <- err
                mu.Unlock()
				return 
			}

			resp, err := http.Post(fmt.Sprintf("%s/uploadRaw?name=%s", uploader, hash), "application/octet-stream", bytes.NewReader(b.RawData()))
			if err != nil {
                mu.Lock()
                errc <- err
                mu.Unlock()
				return
			}

			var value string
			if resp.StatusCode == http.StatusOK {
				var res map[string]interface{}
				json.NewDecoder(resp.Body).Decode(&res)
				value = res["fileId"].(string)
			}

			f := fileInfo{value, len(b.RawData())}
			bf, err := json.Marshal(f)
			if err != nil {
                mu.Lock()
                errc <- err
                mu.Unlock()
				return
			}
			// set
			err = tikv.Puts(b.Cid().Bytes(), []byte(bf))
			if err != nil {
                mu.Lock()
                errc <- err
                mu.Unlock()
				return
			}
		}(b)
	}
    go func() {
        wg.Wait()
        close(errc)
    }()

    var errors []error
    for err := range errc {
        if err != nil {
            errors = append(errors, err)
        }
    }

    if len(errors) > 0 {
        return fmt.Errorf("%d errors occurred during upload: %v", len(errors), errors)
    }

    if s.exchange != nil {
        logger.Debugf("BlockService.BlockAdded %d blocks", len(toput))
        if err := s.exchange.NotifyNewBlocks(ctx, toput...); err != nil {
            logger.Errorf("NotifyNewBlocks: %s", err.Error())
        }
    }
	return nil
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

	kv1, err := tikv.Get(c.Bytes())

	if err == nil {
		var f fileInfo

		err := json.Unmarshal(kv1.V, &f)
		if err != nil {
			return nil, err
		}

		endpoint, err := url.Parse(fmt.Sprintf("%s/cacheFile/%s", uploader, f.FileId))
		if err != nil {
			return nil, err
		}
		fmt.Println(f.FileId, " ", f.Size, c.String())

		rawQuery := endpoint.Query()
		rawQuery.Set("range", fmt.Sprintf("0,%d", f.Size))
		endpoint.RawQuery = rawQuery.Encode()
		fileUrl := endpoint.String()

		fmt.Println(fileUrl)

		resp, err := http.Get(fileUrl)
		if err != nil {
			return nil, err
		}

		bdata, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			//TODO: call api add bandwidth user
			return blocks.NewBlockWithCid(bdata, c)
		}
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
	kv1, err := tikv.Get(c.Bytes())
	if err == nil {
		var f fileInfo

		err := json.Unmarshal(kv1.V, &f)
		if err != nil {
			return nil, err
		}

		endpoint, err := url.Parse(fmt.Sprintf("%s/cacheFile/%s", uploader, f.FileId))
		if err != nil {
			return nil, err
		}
		fmt.Println("getBlocks ", f.FileId, " ", f.Size)

		rawQuery := endpoint.Query()
		rawQuery.Set("range", fmt.Sprintf("0,%d", f.Size))
		endpoint.RawQuery = rawQuery.Encode()
		fileUrl := endpoint.String()

		fmt.Println(fileUrl)

		resp, err := http.Get(fileUrl)
		if err != nil {
			return nil, err
		}

		bdata, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			return blocks.NewBlockWithCid(bdata, c)
		}
	}
	return nil, err
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

func uploadRequest(hash string, rawData []byte) (*http.Response, error) {
    url := fmt.Sprintf("%s/uploadRaw?name=%s", uploader, hash)
    return http.Post(url, "application/octet-stream", bytes.NewReader(rawData))
}

var _ BlockGetter = (*Session)(nil)