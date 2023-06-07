---
description: /api/pinning/pinByHash/
---

# Pin a file to IPFS by its CID

{% swagger method="post" path="" baseUrl="https://api-ipfs.attoaioz.cyou/api/pinning/pinByHash" summary="" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="header" name="pinning_api_key" required="true" %}

{% endswagger-parameter %}

{% swagger-parameter in="header" name="pinning_secret_key" required="true" %}

{% endswagger-parameter %}

{% swagger-response status="200: OK" description="" %}
```
{
    "data": {
        "id": "e5c4456d-5e5c-465d-a02a-3a536fc6e718",
        "file_record_id": "00000000-0000-0000-0000-000000000000",
        "root_hash": "1220cafafc20fc9de1c913769d254ded6be39a12c21bb0028654cd838775aed6325c",
        "cid": "Qmc1135ziMvmFG534i75E8HpJoLqzLzgKBxfjBV9cBsMAs",
        "user_id": "c2452942-5dc5-4a0d-9c38-77a66edcf43a",
        "date_pinned": "2023-06-07T04:16:17.336658824Z",
        "date_unpinned": "0001-01-01T00:00:00Z",
        "pinned": false,
        "is_pin_by_hash": true,
        "is_dir": false,
        "metadata": {
            "name": "name-ipfs-hash"
        },
        "status": "RETRIEVAL"
    },
    "status": "success"
}
```
{% endswagger-response %}
{% endswagger %}

{% tabs %}
{% tab title="cURL" %}
```
curl --location --request POST 'https://api-ipfs.attoaioz.cyou/api/pinning/pinByHash' \
--header 'pinning_api_key: OOC4HF2dGRlgAVzg6vbypg==' \
--header 'pinning_secret_key: 0NFNueE1IKn0bbIMB8cRzG2/JeuIwc0BX/2exij8wco=' \
--header 'Content-Type: application/json' \
--data-raw '{
    "hash_to_pin": "Qmc1135ziMvmFG534i75E8HpJoLqzLzgKBxfjBV9cBsMAs",
    "metadata": {
        "name": "name-ipfs-hash"
    }
}'
```
{% endtab %}

{% tab title="Node.js" %}
```javascript
var axios = require('axios');
var data = JSON.stringify({
  "hash_to_pin": "Qmc1135ziMvmFG534i75E8HpJoLqzLzgKBxfjBV9cBsMAs",
  "metadata": {
    "name": "name-ipfs-hash"
  }
});

var config = {
  method: 'post',
  url: 'https://api-ipfs.attoaioz.cyou/api/pinning/pinByHash',
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

url = "https://api-ipfs.attoaioz.cyou/api/pinning/pinByHash"

payload = json.dumps({
  "hash_to_pin": "Qmc1135ziMvmFG534i75E8HpJoLqzLzgKBxfjBV9cBsMAs",
  "metadata": {
    "name": "name-ipfs-hash"
  }
})
headers = {
  'pinning_api_key': 'OOC4HF2dGRlgAVzg6vbypg==',
  'pinning_secret_key': '0NFNueE1IKn0bbIMB8cRzG2/JeuIwc0BX/2exij8wco=',
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

  url := "https://api-ipfs.attoaioz.cyou/api/pinning/pinByHash"
  method := "POST"

  payload := strings.NewReader(`{
    "hash_to_pin": "Qmc1135ziMvmFG534i75E8HpJoLqzLzgKBxfjBV9cBsMAs",
    "metadata": {
        "name": "name-ipfs-hash"
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
