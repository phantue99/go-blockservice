# Billing

## Retrieves historical usage data

{% swagger method="get" path="" baseUrl="https://api.w3ipfs.storage/api/billing/historyUsage" summary="" %}
{% swagger-description %}
Retrieves historical usage data for a user based on their API key
{% endswagger-description %}

{% swagger-parameter in="header" name="pinning_api_key" required="true" %}
PINNING-API-KEY
{% endswagger-parameter %}

{% swagger-parameter in="query" name="offset" required="false" %}
(default: 0)
{% endswagger-parameter %}

{% swagger-parameter in="header" required="true" name="pinning_secret_key" %}
PINNING-SECRET-KEY
{% endswagger-parameter %}

{% swagger-parameter in="query" name="limit" required="false" %}
(default: 10)
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="" %}
```json
{
    "data": {
        "total_days": number,
        "history_usages": [
            {
                "date": "2023-01-01T11:11:11.111111Z",
                "total_storage": number,
                "total_bandwidth": number,
                "total_amount": "string"
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
curl --location --request GET 'https://api.w3ipfs.storage/api/billing/historyUsage?offset=0&limit=10' \
--header 'pinning_api_key: KEY' \
--header 'pinning_secret_key: SECRET'
```
{% endtab %}

{% tab title="Node.js" %}
```javascript
var axios = require('axios');

var config = {
  method: 'get',
  url: 'https://api.w3ipfs.storage/api/billing/historyUsage?offset=0&limit=10',
  headers: { 
    'pinning_api_key': 'KEY', 
    'pinning_secret_key': 'SECRET'
  }
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

url = "https://api.w3ipfs.storage/api/billing/historyUsage?offset=0&limit=10"

payload = ""
headers = {
  'pinning_api_key': 'KEY',
  'pinning_secret_key': 'SECRET'
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

  url := "https://api.w3ipfs.storage/api/billing/historyUsage?offset=0&limit=10"
  method := "GET"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("pinning_api_key", "KEY")
  req.Header.Add("pinning_secret_key", "SECRET")

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

## Retrieves top up data

{% swagger method="get" path="" baseUrl="https://api.w3ipfs.storage/api/billing/topUp" summary="" %}
{% swagger-description %}
Retrieves top up data for a user based on their API key
{% endswagger-description %}

{% swagger-parameter in="header" name="pinning_api_key" required="true" %}
PINNING-API-KEY
{% endswagger-parameter %}

{% swagger-parameter in="query" name="offset" required="false" %}
(default: 0)
{% endswagger-parameter %}

{% swagger-parameter in="query" name="limit" required="false" %}
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
                "created_at": "2023-01-01T11:11:11.111111Z",
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
curl --location --request GET 'https://api.w3ipfs.storage/api/billing/topUp?offset=0&limit=10' \
--header 'pinning_api_key: KEY' \
--header 'pinning_secret_key: SECRET' \
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
  url: 'https://api.w3ipfs.storage/api/billing/topUp?offset=0&limit=10',
  headers: { 
    'pinning_api_key': 'KEY', 
    'pinning_secret_key': 'SECRET', 
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

url = "https://api.w3ipfs.storage/api/billing/topUp?offset=0&limit=10"

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
  'pinning_api_key': 'KEY',
  'pinning_secret_key': 'SECRET',
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

  url := "https://api.w3ipfs.storage/api/billing/topUp?offset=0&limit=10"
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
  req.Header.Add("pinning_api_key", "KEY")
  req.Header.Add("pinning_secret_key", "SECRET")
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

## Get this month usage data for a user

{% swagger method="get" path="" baseUrl="https://api.w3ipfs.storage/api/billing/thisMonthUsage" summary="" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="query" name="offset" required="false" %}
(default: 0)
{% endswagger-parameter %}

{% swagger-parameter in="header" name="pinning_api_key" required="true" %}
PINNING-API-KEY
{% endswagger-parameter %}

{% swagger-parameter in="query" name="limit" required="false" %}
(default: 10)
{% endswagger-parameter %}

{% swagger-parameter in="header" name="pinning_secret_key" required="true" %}
PINNING-SECRET-KEY
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="" %}
```json
{
    "data": {
        "total_storage": number,
        "storage_usage": number,
        "bandwidth_usage": number,
        "storage_cost": "string",
        "bandwidth_cost": "string"
    },
    "status": "success"
}
```
{% endswagger-response %}
{% endswagger %}

{% tabs %}
{% tab title="cURL" %}
```
curl --location --request GET 'https://api.w3ipfs.storage/api/billing/thisMonthUsage?offset=0&limit=10' \
--header 'pinning_api_key: KEY' \
--header 'pinning_secret_key: SECRET'
```
{% endtab %}

{% tab title="Node.js" %}
```javascript
var axios = require('axios');

var config = {
  method: 'get',
  url: 'https://api.w3ipfs.storage/api/billing/thisMonthUsage?offset=0&limit=10',
  headers: { 
    'pinning_api_key': 'KEY', 
    'pinning_secret_key': 'SECRET'
  }
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

url = "https://api.w3ipfs.storage/api/billing/thisMonthUsage?offset=0&limit=10"

payload={}
headers = {
  'pinning_api_key': 'KEY',
  'pinning_secret_key': 'SECRET'
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
  "net/http"
  "io/ioutil"
)

func main() {

  url := "https://api.w3ipfs.storage/api/billing/thisMonthUsage?offset=0&limit=10"
  method := "GET"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("pinning_api_key", "KEY")
  req.Header.Add("pinning_secret_key", "SECRET")

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
