# Store and mint NFTs using ERC-1155 metadata standards

## Introduction

This document will guide you through the process of using our service to store and mint NFTs using the ERC-1155 metadata standards. With our service, you can easily upload and store your NFT assets, generate ERC-1155 compliant metadata, and store the metadata on IPFS. You will also learn how to mint your own ERC-1155 tokens by providing the IPFS URL of the metadata in the minting transaction.

## Uploading Images, Assets, and Metadata

Before creating the blockchain record for your NFT, you need to store all the off-chain resources that make up the NFT "package." Our service allows you to store the NFT's images and assets, generate ERC-1155 compliant metadata, and link the metadata to the assets using IPFS URIs.

To store your NFT assets and metadata, you can use either the JavaScript client or the HTTP API. Both methods support storing metadata and asset files in a single request and automatically update the metadata to link to the asset files using IPFS URIs.

For the purpose of this documentation, we will provide examples using JavaScript and the HTTP API.

### JavaScript Client Example

The JavaScript client's `store(token)` method is used to store NFT assets and metadata. It takes a single `token` argument, which contains the metadata for your NFT as a JavaScript object.

Here's an example using the JavaScript client:

```javascript
import W3IpfsClient from 'aioz-w3ipfs-sdk';
import fs from 'fs';

// Initialize the client with your API key
let client = new W3IpfsClient("key", "secret-key")

const readableStreamForFile = fs.createReadStream('./test.png');
// Example metadata object
const metadata = {
  name: "My Awesome NFT",
  description: "This is an NFT that represents my creativity as a digital artist!",
  image: null,
  properties: {
    type: "image",
    origins: {
      http: "https://.../",
      ipfs: "ipfs://Qm..."
    },
    authors: [
      { name: "Author" }
    ],
    content: {
      "text/markdown": "Lorem ipsum dolor sit amet, consectetur adipiscing elit..."
    }
  }
}
const options = {
    w3IpfsMetadata: {
        name: "My Awesome NFT",
        keyvalues: {
            customKey: 'customValue',
            customKey2: 'customValue2'
        }
    }
};
// Store the NFT assets and metadata
async function storeNFT() {
  const result = await client.pinNft(metadata, undefined, readableStreamForFile, options)
  console.log('Result: ', result)
}

storeNFT()

```

In this example, we initialize the client with your API key, create a metadata object that conforms to the ERC-1155 standard, and call the `store` method to store the NFT assets and metadata. The method returns the metadata URI, which you can use in the minting process.

## Minting your NFT

Once you have the IPFS URI for your metadata, you can proceed to mint your NFT. The exact process for minting an NFT depends on the blockchain you're using and your specific requirements.

To mint an NFT, you typically need to interact with a smart contract that handles the minting process. You'll need to call a function in the smart contract that assigns a new token ID and sets the metadata URI for the token.

Here's an example of a simple smart contract function for minting ERC-1155 tokens:

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "ERC1155.sol";

contract MyNFT is ERC1155 {
    uint256 public tokenCounter;

    constructor() ERC1155("") {
        tokenCounter = 0;
    }

    function mintNFT(string memory _metadataURI) public {
        tokenCounter++;
        _mint(msg.sender, tokenCounter, 1, "");
        _setURI(tokenCounter, _metadataURI);
    }
}
```

In this example, the `mintNFT` function is called to mint a new token. It increments the `tokenCounter`, mints a token with the new ID, and sets the metadata URI using the `_setURI` function.

When minting your NFT, make sure to provide the properly formatted IPFS URI (e.g., `ipfs://Qm...`) as the `_metadataURI` parameter in the minting transaction.

By following the steps outlined in this documentation, you can easily upload your NFT assets, generate compliant metadata, store the metadata on IPFS, and mint your ERC-1155 tokens.
