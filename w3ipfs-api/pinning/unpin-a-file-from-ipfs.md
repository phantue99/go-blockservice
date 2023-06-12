---
description: /api/pinning/unpin/:pinId
---

# Unpin a file from IPFS

{% swagger method="get" path="" baseUrl="https://api-ipfs.attoaioz.cyou/api/pinning/unpin/:pinId" summary="" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="header" name="pinning_api_key" %}
PINNING-API-KEY
{% endswagger-parameter %}

{% swagger-parameter in="header" name="pinning_secret_key" %}
PINNING-SECRET-KEY
{% endswagger-parameter %}

{% swagger-parameter in="path" required="true" %}
pinId
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="" %}
```json
{
    "data": {
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
        "is_dir": false,
        "metadata": {
            "name": "string"
        },
        "status": "string"
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
--header 'pinning_api_key: KEY' \
--header 'pinning_secret_key: SECRET'
```
{% endtab %}

{% tab title="Node.js" %}
```javascript
var axios = require('axios');

var config = {
  method: 'delete',
  url: 'https://api-ipfs.attoaioz.cyou/api/pinning/unpin/e5c4456d-5e5c-465d-a02a-3a536fc6e718',
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

url = "https://api-ipfs.attoaioz.cyou/api/pinning/unpin/e5c4456d-5e5c-465d-a02a-3a536fc6e718"

payload={}
headers = {
  'pinning_api_key': 'KEY',
  'pinning_secret_key': 'SECRET'
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
