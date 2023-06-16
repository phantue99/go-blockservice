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

{% swagger-parameter in="body" name="name" required="true" %}

{% endswagger-parameter %}

{% swagger-parameter in="body" name="scopes" type="" %}

{% endswagger-parameter %}

{% swagger-response status="200: OK" description="" %}
```json
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

{% tabs %}
{% tab title="cURL" %}
```
curl --location --request POST 'https://api-ipfs.attoaioz.cyou/api/apiKeys/' \
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
{% endtab %}

{% tab title="Node.js" %}
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
  url: 'https://api-ipfs.attoaioz.cyou/apiKeys/',
  headers: { 
    'Authorization': 'Bearer JWT', 
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
{% endtab %}

{% tab title="Python" %}
```python
import requests
import json

url = "https://api-ipfs.attoaioz.cyou/api/apiKeys/"

payload = json.dumps({
  "name": "test-api-key",
  "scopes": {
    "admin": True,
    "data": {
      "pin_list": True,
      "nft_list": True
    },
    "pinning": {
      "unpin": True,
      "pin_by_hash": True,
      "pin_file_to_ipfs": True
    },
    "pin_nft": {
      "unpin_nft": True,
      "pin_nft_to_ipfs": True
    }
  }
})
headers = {
  'Authorization': 'Bearer JWT',
  'Content-Type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```
{% endtab %}

{% tab title="Go" %}
```go
package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "https://api-ipfs.attoaioz.cyou/api/apiKeys/"
  method := "POST"

  payload := strings.NewReader(`{
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
}`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Authorization", "Bearer JWT")
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}
```
{% endtab %}
{% endtabs %}

*
