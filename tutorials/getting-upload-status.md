# Getting upload status

When using the AIOZ W3IPFS pinning service for uploading and storing data, you can check the status of your upload using the following methods.

1. [File listing page](https://w3ipfs.storage/dashboard/ipfs-files): AIOZ W3IPFS provides a file listing page where you can see a list of the files you've uploaded.
   * `Pinning`: This status indicates that your upload is in the queue and awaiting processing by the pinning service.
   * `Pinned`: The pinned status confirms that your data has been successfully stored and pinned on the IPFS network.
   * `Failed`: This status indicates that the data upload has failed or the graph of data is incomplete.
2. JavaScript client's check method: The [AIOZ W3IPFS JavaScript client](../sdk.md#w3ipfs-api-node-sdk) provides a check method that you can use to retrieve information about the upload status.
3. HTTP API: AIOZ W3IPFS offers an HTTP API that allows you to send a GET request to the [`/pinning/{pinId}`](../w3ipfs-api/pinning.md#get-pin-details-by-cid) endpoint to fetch the status of your upload. Replace `{pinId}` with the ID of your uploaded data. The response will include the status information, including whether the upload is complete or if it has failed.

{% hint style="warning" %}
If you have a large upload with numerous blocks, such as a directory containing a significant number of files (e.g., a directory for your 10K NFT drop), the AIOZ W3IPFS pinning service may show the status as `pinning` indefinitely. This occurs when the upload is too large for immediate validation of its completeness as a graph.
{% endhint %}
