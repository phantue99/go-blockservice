---
description: /api/pinning/pins
---

# Retrieve pins associated with a user

{% swagger method="get" path="" baseUrl="https://api-ipfs.attoaioz.cyou/api/pinning/pins/" summary="" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="header" name="pinning_api_key" required="true" %}
PINNING-API-KEY
{% endswagger-parameter %}

{% swagger-parameter in="header" name="pinning_secret_key" required="true" %}
PINNING-SECRET-KEY
{% endswagger-parameter %}

{% swagger-parameter in="query" name="offset" %}
(default 0)
{% endswagger-parameter %}

{% swagger-parameter in="query" name="limit" %}
(default 10)
{% endswagger-parameter %}

{% swagger-parameter in="query" name="pinned" %}
Filter by pinned status (options: all, true, false) (default all)
{% endswagger-parameter %}

{% swagger-parameter in="query" name="sortBy" %}
Field to sort by (options: created_at, size, name). Defaults to created_at.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="sortOrder" %}
Sort direction (options: ASC, DESC). Defaults to DESC.
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="" %}
```json
{
    "data": {
        "totals": {
            "files": number,
            "size": number
        },
        "pins": [
            {
                "id": "string",
                "file_record_id": "string",
                "root_hash": "string",
                "cid": "string",
                "size": number,
                "user_id": "string",
                "date_pinned": "2023-01-01T11:11:11.111111Z",
                "date_unpinned": "2023-11-11T11:11:11.111111Z",
                "pinned": true,
                "is_pin_by_hash": true,
                "sub_hash_status": "string",
                "is_dir": false,
                "metadata": {
                    "name": "string"
                },
                "status": "string"
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
curl --location --request GET 'https://api-ipfs.attoaioz.cyou/api/pinning/pins/?offset=0&limit=10&pinned=true&sortBy=name&sortOrder=ASC' \
--header 'pinning_api_key: OOC4HF2dGRlgAVzg6vbypg==' \
--header 'pinning_secret_key: 0NFNueE1IKn0bbIMB8cRzG2/JeuIwc0BX/2exij8wco='
```
{% endtab %}

{% tab title="Node.js" %}
```javascript
var axios = require('axios');

var config = {
  method: 'get',
  url: 'https://api-ipfs.attoaioz.cyou/api/pinning/pins/?offset=0&limit=10&pinned=true&sortBy=name&sortOrder=ASC',
  headers: { 
    'pinning_api_key': 'OOC4HF2dGRlgAVzg6vbypg==', 
    'pinning_secret_key': '0NFNueE1IKn0bbIMB8cRzG2/JeuIwc0BX/2exij8wco='
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

url = "https://api-ipfs.attoaioz.cyou/api/pinning/pins/?offset=0&limit=10&pinned=true&sortBy=name&sortOrder=ASC"

payload={}
headers = {
  'pinning_api_key': 'OOC4HF2dGRlgAVzg6vbypg==',
  'pinning_secret_key': '0NFNueE1IKn0bbIMB8cRzG2/JeuIwc0BX/2exij8wco='
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

  url := "https://api-ipfs.attoaioz.cyou/api/pinning/pins/?offset=0&limit=10&pinned=true&sortBy=name&sortOrder=ASC"
  method := "GET"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("pinning_api_key", "OOC4HF2dGRlgAVzg6vbypg==")
  req.Header.Add("pinning_secret_key", "0NFNueE1IKn0bbIMB8cRzG2/JeuIwc0BX/2exij8wco=")

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
}g
```
{% endtab %}
{% endtabs %}
