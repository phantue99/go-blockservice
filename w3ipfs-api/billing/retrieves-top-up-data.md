---
description: /api/billing/topUp/
---

# Retrieves top up data

{% swagger method="get" path="" baseUrl="https://api-ipfs.attoaioz.cyou/api/billing/topUp" summary="" %}
{% swagger-description %}
Retrieves top up data for a user based on their API key
{% endswagger-description %}

{% swagger-parameter in="header" name="pinning_api_key" required="true" %}
PINNING-API-KEY
{% endswagger-parameter %}

{% swagger-parameter in="query" name="offset" %}
(default: 0)
{% endswagger-parameter %}

{% swagger-parameter in="query" name="limit" %}
(default: 10)
{% endswagger-parameter %}

{% swagger-parameter in="header" name="pinning_secret_key" required="true" %}
PINNING-SECRET-KEY
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="" %}
```json
{
    "data": {
        "totals": number,
        "top_up_usages": [
            {
                "cosmos_tx_hash": "string",
                "event_index": number,
                "evm_tx_hash": "string",
                "id": "string",
                "sender": "string",
                "recipient": "string",
                "block_number": number,
                "status": true,
                "amount": {
                    "denom": "string",
                    "amount": "string"
                },
                "total_amount": "string",
                "created_at": "2023-01-01T11:11:11.111111Z"",
                "updated_at": "2023-01-01T11:11:11.111111Z"
            }
        ]
    },
    "status": "success"
}
```
{% endswagger-response %}
{% endswagger %}

{% tabs %}
{% tab title="cURL" %}
```
curl --location --request GET 'https://api-ipfs.attoaioz.cyou/api/billing/topUp?offset=0&limit=10' \
--header 'pinning_api_key: OOC4HF2dGRlgAVzg6vbypg==' \
--header 'pinning_secret_key: 0NFNueE1IKn0bbIMB8cRzG2/JeuIwc0BX/2exij8wco=' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "test-api-key3",
    "scopes": {
        "data": {
            "pinList": true
        },
        "pinning": {
            "unpin": true
        }
    }
}'
```
{% endtab %}

{% tab title="Node.js" %}
```javascript
var axios = require('axios');
var data = JSON.stringify({
  "name": "test-api-key3",
  "scopes": {
    "data": {
      "pinList": true
    },
    "pinning": {
      "unpin": true
    }
  }
});

var config = {
  method: 'get',
  url: 'https://api-ipfs.attoaioz.cyou/api/billing/topUp?offset=0&limit=10',
  headers: { 
    'pinning_api_key': 'OOC4HF2dGRlgAVzg6vbypg==', 
    'pinning_secret_key': '0NFNueE1IKn0bbIMB8cRzG2/JeuIwc0BX/2exij8wco=', 
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

url = "https://api-ipfs.attoaioz.cyou/api/billing/topUp?offset=0&limit=10"

payload = json.dumps({
  "name": "test-api-key3",
  "scopes": {
    "data": {
      "pinList": True
    },
    "pinning": {
      "unpin": True
    }
  }
})
headers = {
  'pinning_api_key': 'OOC4HF2dGRlgAVzg6vbypg==',
  'pinning_secret_key': '0NFNueE1IKn0bbIMB8cRzG2/JeuIwc0BX/2exij8wco=',
  'Content-Type': 'application/json'
}

response = requests.request("GET", url, headers=headers, data=payload)

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

  url := "https://api-ipfs.attoaioz.cyou/api/billing/topUp?offset=0&limit=10"
  method := "GET"

  payload := strings.NewReader(`{
    "name": "test-api-key3",
    "scopes": {
        "data": {
            "pinList": true
        },
        "pinning": {
            "unpin": true
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
  req.Header.Add("pinning_api_key", "OOC4HF2dGRlgAVzg6vbypg==")
  req.Header.Add("pinning_secret_key", "0NFNueE1IKn0bbIMB8cRzG2/JeuIwc0BX/2exij8wco=")
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
