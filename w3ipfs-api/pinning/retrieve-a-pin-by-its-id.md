---
description: /api/pinning/:pinId
---

# Retrieve a pin by its ID

{% swagger method="get" path="" baseUrl="https://api-ipfs.attoaioz.cyou/api/pinning/:pinId" summary="" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="header" name="pinning_api_key" required="true" %}
PINNING-API-KEY
{% endswagger-parameter %}

{% swagger-parameter in="header" name="pinning_secret_key" required="true" %}
PINNING-SECRET-KEY
{% endswagger-parameter %}

{% swagger-parameter in="path" required="true" %}
pinId
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="" %}
```
{
    "data": {
        "id": "3c3fec2a-ca65-4b8e-bcf3-c8e2ceaa23d2",
        "file_record_id": "19164af7-8088-407d-8b4f-5bb3bae3541a",
        "root_hash": "1220ce95a46e2063878d025d68ce543a205eee90e5b31a659cb8c7c108fb3b133414",
        "cid": "bafkreigoswsg4iddq6gqexlizzkduic652iolmy2mwolrr6bbd5twezucq",
        "size": 50922,
        "user_id": "c2452942-5dc5-4a0d-9c38-77a66edcf43a",
        "date_pinned": "2023-06-07T03:52:39.331369Z",
        "date_unpinned": "0001-01-01T00:00:00Z",
        "pinned": true,
        "is_pin_by_hash": false,
        "sub_hash_status": "DONE",
        "is_dir": false,
        "metadata": {
            "name": "test.png",
            "type": "image/png"
        },
        "status": "DONE"
    },
    "status": "success"
}
```
{% endswagger-response %}
{% endswagger %}

{% tabs %}
{% tab title="cURL" %}
```
curl --location --request GET 'https://api-ipfs.attoaioz.cyou/api/pinning/3c3fec2a-ca65-4b8e-bcf3-c8e2ceaa23d2' \
--header 'pinning_api_key: OOC4HF2dGRlgAVzg6vbypg==' \
--header 'pinning_secret_key: 0NFNueE1IKn0bbIMB8cRzG2/JeuIwc0BX/2exij8wco=' \
--data-raw ''
```
{% endtab %}

{% tab title="Node.js" %}
```javascript
var axios = require('axios');
var data = '';

var config = {
  method: 'get',
  url: 'https://api-ipfs.attoaioz.cyou/api/pinning/3c3fec2a-ca65-4b8e-bcf3-c8e2ceaa23d2',
  headers: { 
    'pinning_api_key': 'OOC4HF2dGRlgAVzg6vbypg==', 
    'pinning_secret_key': '0NFNueE1IKn0bbIMB8cRzG2/JeuIwc0BX/2exij8wco='
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

url = "https://api-ipfs.attoaioz.cyou/api/pinning/3c3fec2a-ca65-4b8e-bcf3-c8e2ceaa23d2"

payload = ""
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
  "strings"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "https://api-ipfs.attoaioz.cyou/api/pinning/3c3fec2a-ca65-4b8e-bcf3-c8e2ceaa23d2"
  method := "GET"

  payload := strings.NewReader(``)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

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
}
```
{% endtab %}
{% endtabs %}
