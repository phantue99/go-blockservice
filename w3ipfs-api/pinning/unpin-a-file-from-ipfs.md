---
description: /api/pinning/unpin/:pinId
---

# Unpin a file from IPFS

{% swagger method="get" path="" baseUrl="https://api-ipfs.attoaioz.cyou/api/pinning/unpin/:pinId" summary="" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="header" name="pinning_api_key" %}

{% endswagger-parameter %}

{% swagger-parameter in="header" name="pinning_secret_key" %}

{% endswagger-parameter %}

{% swagger-parameter in="path" required="true" %}
pinId
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="" %}
```
{
    "data": {
        "id": "e5c4456d-5e5c-465d-a02a-3a536fc6e718",
        "file_record_id": "19164af7-8088-407d-8b4f-5bb3bae3541a",
        "root_hash": "1220cafafc20fc9de1c913769d254ded6be39a12c21bb0028654cd838775aed6325c",
        "cid": "Qmc1135ziMvmFG534i75E8HpJoLqzLzgKBxfjBV9cBsMAs",
        "size": 2885325,
        "user_id": "c2452942-5dc5-4a0d-9c38-77a66edcf43a",
        "date_pinned": "2023-06-07T04:16:39.641031Z",
        "date_unpinned": "0001-01-01T00:00:00Z",
        "pinned": true,
        "is_pin_by_hash": true,
        "is_dir": false,
        "metadata": {
            "name": "name-ipfs-hash"
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
curl --location --request DELETE 'https://api-ipfs.attoaioz.cyou/api/pinning/unpin/:pinId
--header 'pinning_api_key: OOC4HF2dGRlgAVzg6vbypg==' \
--header 'pinning_secret_key: 0NFNueE1IKn0bbIMB8cRzG2/JeuIwc0BX/2exij8wco='
```
{% endtab %}

{% tab title="Node.js" %}
```javascript
var axios = require('axios');

var config = {
  method: 'delete',
  url: 'https://api-ipfs.attoaioz.cyou/api/pinning/unpin/e5c4456d-5e5c-465d-a02a-3a536fc6e718',
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

url = "https://api-ipfs.attoaioz.cyou/api/pinning/unpin/e5c4456d-5e5c-465d-a02a-3a536fc6e718"

payload={}
headers = {
  'pinning_api_key': 'OOC4HF2dGRlgAVzg6vbypg==',
  'pinning_secret_key': '0NFNueE1IKn0bbIMB8cRzG2/JeuIwc0BX/2exij8wco='
}

response = requests.request("DELETE", url, headers=headers, data=payload)

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

  url := "https://api-ipfs.attoaioz.cyou/api/pinning/unpin/e5c4456d-5e5c-465d-a02a-3a536fc6e718"
  method := "DELETE"

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
}
```
{% endtab %}
{% endtabs %}
