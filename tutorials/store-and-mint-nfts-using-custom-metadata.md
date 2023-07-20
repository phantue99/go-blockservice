# Store and mint NFTs using custom metadata

## Introduction&#x20;

The W3IPFS Pinning Service is a service that allows users to store and mint NFTs using custom metadata. It leverages the power of the InterPlanetary File System (IPFS) for decentralized and resilient storage of NFT assets and metadata. This document provides an overview of how to use the W3IPFS Pinning Service to upload and store your NFT assets, prepare custom metadata, and mint your own NFTs.

## Uploading and Storing NFT Assets

To create the blockchain record for your NFT, you first need to store all the off-chain resources that make up the NFT "package". The W3IPFS Pinning Service provides methods to store arbitrary files for your NFT assets and metadata.

### Storing Asset Files&#x20;

* `pinFile`: Accepts a File object and returns the Content Identifier (CID) of the uploaded file. Note that the original filename and content type information are not preserved.
* `pinFolder`: Accepts multiple File objects and creates an IPFS directory listing, allowing you to link to files using human-readable names as the "path" component of IPFS URIs. It returns the CID of the directory listing.

### Preparing Custom Metadata

Once you have stored all your assets, you can update your metadata to include IPFS URIs pointing to the images and other files associated with your NFT. Follow these steps to prepare your custom metadata:

* Constructing IPFS URIs: For each uploaded file, create an IPFS URI of the form `ipfs://<CID>/<filename>`. This ensures that any IPFS-compatible browser can retrieve the data directly using these URLs.
* Including HTTP Gateway URLs (optional): While IPFS URIs are recommended, you can include HTTP gateway URLs as an optimization or fallback for compatibility with browsers that don't support IPFS natively. Be aware of the tradeoffs involved in using HTTP gateways.
* Serialize Metadata: Serialize your metadata into a file, typically in JSON format. Include the IPFS URIs or HTTP gateway URLs for your assets within the metadata.
* Upload Metadata: upload metadata file follow [here](../w3ipfs-api/pinning.md#pin-files-or-directory)

## Minting Your NFTs

With the metadata stored using the W3IPFS Pinning Service, you can now mint tokens using the blockchain platform of your choice. The process of minting depends on the specific blockchain, development language, contract, and standards you are targeting.

Important Considerations:

* Avoid storing HTTP gateway links in smart contracts or blockchain records. Store the `ipfs://` URI as the canonical link and convert it to a gateway link when displaying the NFT on the web.
* Prefer using IPFS URIs instead of raw CIDs or hashes for linking to IPFS data. URIs of the form `ipfs://<CID>/<path>` provide flexibility and human-friendliness.

Conclusion:

The W3IPFS Pinning Service empowers users to store and mint NFTs using custom metadata. By leveraging IPFS for decentralized storage and retrieval, the service ensures the resilience and integrity of NFT assets. Follow the outlined steps to upload and store your assets, prepare custom metadata, and mint your own NFTs using the W3IPFS Pinning Service.
