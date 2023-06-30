# Pinning Services API

## Pinning Services API

The Pinning Services API is a standardized specification for developers who want to build applications on top of IPFS. It allows seamless integration with a pinning service without requiring knowledge of the service's unique API.

### Endpoint

W3IPFS users looking to utilize the IPFS Pinning Services API can do so from the dedicated API endpoint: [`https://api-ipfs.attoaioz.cyou/api/psa/`](https://api-ipfs.attoaioz.cyou/api/psa/)

## Authentication

To authenticate with W3IPFS through the Pinning Services API spec, you'll need to obtain an `accessToken`. To obtain an access token for authentication with the W3ipfs Pinning Services API, you can use Postman to make a POST request to the following API endpoint:

{% swagger method="post" path="" baseUrl="https://api-ipfs.attoaioz.cyou/api/auth/login" summary="" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="body" name="email" required="true" %}
Your actual email address
{% endswagger-parameter %}

{% swagger-parameter in="body" name="password" required="true" %}
Your password
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="" %}
```
{
    "access_token": "JWT",
    "refresh_token": "JWT",
    "status": "success"
}
```
{% endswagger-response %}
{% endswagger %}

## Using the IPFS CLI

To use the W3IPFS Pinning Service with the IPFS CLI, you can follow these steps:

1. Add W3IPFS credentials: Before you can start pinning to the W3IPFS Pinning Service, you need to add your authentication credentials. Replace `YOUR_AUTH_JWT` in the following command with your actual API key.

```
ipfs pin remote service add w3ipfs https://api-ipfs.attoaioz.cyou/api/psa/ YOUR_AUTH_JWT
```

2. Pin a CID to W3IPFS: Once you have added the service, you can pin a CID to W3IPFS using the `ipfs pin remote add` command. Replace `YOUR_CID` with the CID you want to pin.

```
ipfs pin remote add --service=w3ipfs --name=<human-readable-name> YOUR_CID
```

3. List successful pins: To see a list of successful pins on W3IPFS, you can use the following command:

```
ipfs pin remote ls --service=w3ipfs
```

4. List pending pins: If you want to check the status of pending pins on W3IPFS, you can use the following command:

```
ipfs pin remote ls --service=w3ipfs --status=queued,pinning,failed
```

5. Additional commands and help: For more commands and general help with the IPFS remote pinning feature, you can use the following command:

```
ipfs pin remote --help
```
