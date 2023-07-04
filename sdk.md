# ✊ SDK

## W3IPFS API Node SDK

The official[ ](w3ipfs-api/)[**W3IPFS API**](w3ipfs-api/) Node.js SDK. This is the easiest way to start developing with W3IPFS.&#x20;

{% embed url="https://www.npmjs.com/package/aioz-w3ipfs-sdk" %}

## W3IPFS PINNING CLI

This application provides a Command Line Interface for Pinning APIs.

{% embed url="https://10.0.0.50/tue.phan/pinning-service-cli" %}

### Command line login

`pinning login --key "your_key" --secret "your_secret"`

### Example usage

**You can pin a file**\
`pinning pin --file afile.txt`

**You can pin a whole directory**\
`pinning pin --file ../some/where`

**You can choose add metadata for your own usage**\
`pinning pin --file afile.txt --keyvalue key1:value1 --keyvalue key2:value2`

**Add a hash to be pinned**\
`pinning pin --hash QmdYTBNig2d4dQd5o1LXM3NHbCYA7168NpN5R9m44vDj88 --keyvalue key1:value1 --keyvalue key2:value2`

**Get pin by id**\
`pinning get-pin --id=00000000-0000-0000-0000-000000000000`

**Or list all your pins**\
`pinning list-pins`

**Or your custom**\
`pinning list-pins --pinned=true --sortBy=created_at --sortOrder=DESC --limit=10 --offset=0 --keyvalue=key1:value1 --keyvalue=key2:value2`

**And finally unpin a hash by its ID**\
`pinning unpin --id=00000000-0000-0000-0000-000000000000`

