# 🤝 Limit

## W3IPFS Limits

The W3IPFS platform currently has the following rate limits in place:

#### API Rate Limits

The W3IPFS API currently has a rate limit of 180 requests per minute to our API.&#x20;

#### Exceptions

The following API calls have increased rate limits:

#### Recommendations

For users with traffic that can be burst-heavy, we recommend using a task scheduler or queue-based upload approach that limits requests to the allowed amount.&#x20;

### Public Gateway Rate Limits

The W3IPFS public IPFS gateway is meant for testing purposes or very low volume retrieval and should not be used in production scenarios. It currently has the following rate limits:

* Each CID has a global rate limit of 15 requests per minute (this is across all IP addresses)
* Each IP address has a rate limit of 200 requests per minute

### Dedicated Gateway Rate Limits

At this time there are currently no rate limits for users retrieving content from a dedicated gateway.&#x20;
