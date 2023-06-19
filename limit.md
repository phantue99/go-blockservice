# 🤝 Limit

## W3IPFS Limits

The W3IPFS platform enforces the following rate limits for its API:

#### API Rate Limits

The W3IPFS API has a rate limit of 180 requests per minute.

#### Exceptions

The following API calls have increased rate limits:

* Endpoints under the `/users/` and `/apiKeys/` the route has a rate limit of 60 requests per minute
* The Pinning Services API endpoint for listing content has a rate limit of 60 requests per minute

#### Recommendations

For users with burst-heavy traffic patterns, it is advisable to implement a task scheduler or a queue-based upload approach. This approach helps ensure that the number of requests made to the API stays within the allowed limit.

### Public and Dedicated Gateway Rate Limits

The W3IPFS gateway provides both public and dedicated gateways for retrieving content from the IPFS network. Please note the following rate limits for each gateway type:

#### Public Gateway Rate Limits:

The public IPFS gateway is primarily intended for testing purposes or scenarios with very low retrieval volumes. It is not suitable for production use. The current rate limits for the public gateway are as follows:

* Each CID is subject to a global rate limit of 15 requests per minute. This limit applies across all IP addresses. It ensures fair usage and prevents excessive strain on the public gateway.
* Each IP address accessing the public gateway has a rate limit of 200 requests per minute. This limit helps distribute the load and maintain the gateway's performance for all users.

#### Dedicated Gateway Rate Limits

For users who require higher reliability and performance, W3IPFS offers dedicated gateways. These gateways do not impose any specific rate limits for retrieving content. However, it is important to use dedicated gateways responsibly and avoid excessive requests that could potentially affect the gateway's performance for other users.&#x20;

Please keep in mind that even though dedicated gateways do not have predefined rate limits, it is still essential to use them responsibly and avoid any abusive or disruptive behavior.
