# Retrieve NFT data from IPFS

## Finding the IPFS address for your NFT's metadata

To retrieve the off-chain metadata for your NFT, you need to locate its IPFS address. The IPFS address is typically recorded in the blockchain entry for the token and can be found on various NFT marketplaces and explorer sites. Here's a guide on how to find the IPFS address for your NFT's metadata:

1.  NFT Marketplaces and Explorer Sites: Many NFT marketplaces and explorer sites provide the IPFS address for a token's metadata. For example, on OpenSea, you can find the metadata link under the "Details" view, which includes a link to the "Frozen" metadata stored on IPFS. The metadata link will typically resemble an IPFS gateway URL.

    Example: Metadata Link: [https://ipfs.io/ipfs/bafkreigfvngoydofemwj5x5ioqsaqarvlprzgxinkcv3am3jpv2sysqobi](https://ipfs.io/ipfs/bafkreigfvngoydofemwj5x5ioqsaqarvlprzgxinkcv3am3jpv2sysqobi)

    In the example above, the IPFS gateway URL is "[https://ipfs.io](https://ipfs.io)," and the Content Identifier (CID) is the random-looking string after "/ipfs/" in the link.

    Note: The gateway host and IPFS address format may vary depending on the NFT creation method and the platform you're using.
2.  Block Explorer and Blockchain Directly: If your marketplace or wallet does not display the original metadata URI, you can try using a block explorer to directly consult the blockchain. The process will depend on the blockchain platform and the smart contract standard used to mint your NFT.

    Example: If your NFT was minted on the Ethereum network and follows the ERC-1155 standard, you can interact with the smart contract and call the `tokenURI()` function to retrieve the metadata URI directly.

<figure><img src="../.gitbook/assets/Neko-NEKO-Token-Tracker-Etherscan.png" alt=""><figcaption></figcaption></figure>

Remember, the process of finding the IPFS address may vary depending on the platform, blockchain, and standards used for your specific NFT. Always refer to the documentation or resources provided by the marketplace, wallet, or platform where you obtained or minted the NFT for more specific instructions.

## Option: Retrieve using IPFS HTTP gateways

