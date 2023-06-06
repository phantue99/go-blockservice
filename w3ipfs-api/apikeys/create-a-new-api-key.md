---
description: /api/apiKeys/
---

# Generate W3IPFS API key

This endpoint is used to programmatically generate W3IPFS API keys. This endpoint can only be called by using an "Admin" key. When generating new keys, specific scopes and limits can be implemented.&#x20;

> **Make sure to record the API Secret as they will not be accessible again.**

## Generating an API Key

The request body when generating a W3IPFS API key will look like this:&#x20;

```
{
    name: (A name for your new key for easy reference - Required),
    scopes: {
        admin: boolean,
        data: {
            pin_list: boolean,
            nft_list: boolean
        },
        pinning: {
            unpin: boolean,
            pin_by_hash: boolean,
            pin_file_to_ipfs: boolean
        },
        pin_nft:{
            unpin_nft: boolean,
            pin_nft_to_ipfs: boolean
        }
    }
}
```

Notice the `name` is required. When setting the permissions, it is necessary to include all properties and sub-properties unless you are creating an admin key. If you are creating an admin key, the sub-properties can be omitted.

For example, this would be a simplified body for admin key generation:&#x20;

```
{
    name: "Admin api key",
    scopes: {
        admin: true
    }
}
```

{% swagger method="post" path="" baseUrl="https://api-ipfs.attoaioz.cyou/api/apiKeys/" summary="" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="header" name="pinning_api_key" required="true" %}
PINNING-API-KEY
{% endswagger-parameter %}

{% swagger-parameter in="header" name="pinning_secret_key" required="true" %}
PINNING-SECRET-KEY
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="" %}
```
{
    "data": {
        "name": "API KEY",
        "api_key": "KEY",
        "secret_key": "SECRET"
    },
    "status": "success"
}
```
{% endswagger-response %}
{% endswagger %}

{% code title="" %}
```url
curl --location --request POST '0.0.0.0:8000/api/apiKeys/' \
--header 'Authorization: Bearer JWT' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "test-api-key",
    "scopes": {
        "admin": true,
        "data": {
            "pin_list": true,
            "nft_list":true
        },
        "pinning": {
            "unpin": true,
            "pin_by_hash": true,
            "pin_file_to_ipfs": true
        },
        "pin_nft":{
            "unpin_nft": true,
            "pin_nft_to_ipfs": true
        }
    }
}'
```
{% endcode %}

{% code title="Node.js" overflow="wrap" lineNumbers="true" fullWidth="false" %}
```javascript
var axios = require('axios');
var data = JSON.stringify({
  "name": "test-api-key",
  "scopes": {
    "admin": true,
    "data": {
      "pin_list": true,
      "nft_list": true
    },
    "pinning": {
      "unpin": true,
      "pin_by_hash": true,
      "pin_file_to_ipfs": true
    },
    "pin_nft": {
      "unpin_nft": true,
      "pin_nft_to_ipfs": true
    }
  }
});

var config = {
  method: 'post',
  url: '0.0.0.0:8000/api/apiKeys/',
  headers: { 
    'Authorization': 'Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODYwNDI0ODcsImlhdCI6MTY4NjAzNTI4NywibmJmIjoxNjg2MDM1Mjg3LCJzdWIiOiIwNzE0ZjM1Yi0wOGFhLTRjMmItOWRmNC1kZGFhMjc2ZDRlZDAifQ.dEZm_JcmojJQynZ-5Vc9ryifbpSNTrOn2PyRWw-aN_cudKFfreVUhIY17EewoKHzy88eSuNGs5oRdYkqoHaBzQ3jV6x8rytnxJKGr5QzKri7yALPRQu4tSHJtuxTwa_Ypi_ISS1S2CEfw43c5-luzpT2Cn0i45C6UjbZdPGd0Mw', 
    'Content-Type': 'application/json'
  },
  data : data
};

axios(config)
.then(function (response) {
  console.log(JSON.stringify(response.data));
})
.catch(function (error) {
  console.log(error);
});
```
{% endcode %}

```
// Some code
```

<img alt="" class="gitbook-drawing">

<table data-view="cards"><thead><tr><th></th><th></th><th></th></tr></thead><tbody><tr><td></td><td></td><td></td></tr><tr><td></td><td></td><td></td></tr><tr><td></td><td></td><td></td></tr></tbody></table>

1.
