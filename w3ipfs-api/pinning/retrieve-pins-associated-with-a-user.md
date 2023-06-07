---
description: /api/pinning/pins
---

# Retrieve pins associated with a user

{% swagger method="get" path="" baseUrl="https://api-ipfs.attoaioz.cyou/api/pinning/pins/" summary="" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="header" name="pinning_api_key" required="true" %}

{% endswagger-parameter %}

{% swagger-parameter in="header" name="pinning_secret_key" required="true" %}

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
```
{
    "data": {
        "totals": {
            "files": 100,
            "size": 100000000000
        },
        "pins": [
            {
                "id": "2cfc0003-70d8-4475-a994-3e141e5590ce",
                "file_record_id": "19164af7-8088-407d-8b4f-5bb3bae3541a",
                "root_hash": "1220f7ee5e4ae2c76c8c2ef468034248f3fdb7ee0a8926e9838d987d57c43c9c030c",
                "cid": "bafybeifjcleklgkflfbwgvc2ltveyk3o65kbb3ezkskjl7okfhxmzu5sya",
                "size": 14150060,
                "user_id": "c2452942-5dc5-4a0d-9c38-77a66edcf43a",
                "date_pinned": "2023-05-23T08:29:43.264508Z",
                "date_unpinned": "0001-01-01T00:00:00Z",
                "pinned": true,
                "is_pin_by_hash": false,
                "sub_hash_status": "DONE",
                "is_dir": false,
                "metadata": {
                    "name": "test"
                },
                "status": "DONE"
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
