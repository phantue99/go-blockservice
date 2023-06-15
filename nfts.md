# 🥳 NFTs

Welcome to the NFTs page, where you can create and manage your non-fungible tokens (NFTs). NFTs represent unique assets that can be owned, bought, and sold on the blockchain. This documentation will guide you through the process of creating NFTs using our platform.

## Why Should NFT Assets Be Hosted on IPFS

Hosting NFT assets on IPFS (InterPlanetary File System) offers several advantages and benefits. Here are some reasons why you should consider using IPFS for hosting your NFT assets:

1. **Decentralization:** IPFS is a decentralized file storage system that provides a distributed network for storing and retrieving data. Unlike traditional centralized servers, IPFS allows NFT assets to be stored across multiple nodes in the network, making it more resilient to failures and censorship.
2. **Immutable Content:** IPFS uses content addressing, which means that each file is identified by a unique hash. Once an NFT asset is uploaded and pinned to IPFS, its content is permanently associated with the hash, ensuring the immutability of the asset. This is crucial for the integrity and authenticity of NFTs.
3. **Content Addressing:** IPFS uses cryptographic hashes to address and retrieve content. This means that the content of an NFT asset is uniquely identified by its hash, regardless of its location or storage provider. This enables easy reference and retrieval of NFT assets by their content hash, ensuring that the asset remains accessible even if the hosting location changes.
4. **Efficient Distribution:** IPFS leverages a distributed peer-to-peer network, allowing NFT assets to be distributed efficiently across multiple nodes. When an NFT asset is requested, it can be fetched from the nearest available node, reducing latency and improving performance.
5. **Reducing Costs:** By hosting NFT assets on IPFS, you can potentially reduce the costs associated with traditional centralized hosting services. IPFS allows you to leverage the network's distributed nature, eliminating the need for costly infrastructure and centralized servers.
6. **Interoperability:** IPFS is designed to be protocol-agnostic and can be integrated with various blockchain platforms and applications. This interoperability allows NFT assets hosted on IPFS to be easily referenced and accessed by different blockchain networks and applications, enhancing their compatibility and potential for widespread adoption.
7. **Community and Ecosystem:** IPFS has a vibrant and active community of developers and contributors who are continuously working on improving the protocol and building tools and services around it. By hosting NFT assets on IPFS, you become part of this growing ecosystem and benefit from the collective knowledge and support of the community.

In summary, hosting NFT assets on IPFS provides decentralization, immutability, efficient distribution, reduced costs, interoperability, and access to a thriving community and ecosystem. These advantages make IPFS an excellent choice for ensuring the long-term availability, integrity, and accessibility of your NFT assets.

## Asset and Metadata

1. Asset:
   * An NFT asset can be either a file or a CID (Content Identifier) string.
   * If you choose to use a file as the asset, you can upload it in the form provided on the page.
   * Alternatively, you can provide a CID string as a query parameter to reference the asset stored on IPFS.
2. Metadata:
   * Each NFT requires metadata that provides additional information about the asset.
   * The metadata is represented in a JSON file format.
   * The JSON file should include the following fields:
     * "name": The name of the NFT.
     * "description": A description of the NFT.
     * "properties": Additional properties or attributes associated with the NFT.

## Creating an NFT

1. **Uploading Asset:**
   * If your NFT asset is a file, use the provided form to upload the file.
   * Ensure that the file meets the supported file format requirements.
2. **Adding Metadata:**
   * Upload a JSON file containing the required metadata for the NFT.
   * The JSON file should adhere to the structure mentioned earlier.
3. **Creating the NFT:**
   * Click on the "Create NFT" button to initiate the creation process.
   * The system will process the asset and metadata to mint a unique NFT on the blockchain.
4. **Faster NFT Creation using IPFS:**
   * To expedite the NFT creation process, we support directly creating NFTs from assets stored on IPFS.
   * If you already have the asset CID, you can provide it as a query parameter to create the NFT faster.
5. **Custom Metadata:**
   * If you require additional custom metadata for your NFT, you can create a separate JSON file.
   * Pin the JSON file to IPFS using the usual method and reference the CID in the main metadata JSON.
