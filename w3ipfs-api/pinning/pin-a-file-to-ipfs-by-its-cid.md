---
description: /api/pinning/pinByHash/
---

# Pin a file to IPFS by its CID

The request body when pin a file by CID will look like this:&#x20;

```
{
    hash_to_pin: CID,
    metadata: {
        name: string,
        keyvalues: {
            key1: value1,
            key2: value2
        }
    }
}
```

{% swagger method="post" path="" baseUrl="https://api-ipfs.attoaioz.cyou/api/pinning/pinByHash" summary="" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="header" name="pinning_api_key" required="true" %}
PINNING-API-KEY
{% endswagger-parameter %}

{% swagger-parameter in="header" name="pinning_secret_key" required="true" %}
PINNING-SECRET-KEY
{% endswagger-parameter %}

{% swagger-parameter in="body" name="hash_to_pin" required="true" %}
CID
{% endswagger-parameter %}

{% swagger-parameter in="body" name="metadata" %}

{% endswagger-parameter %}

{% swagger-response status="200: OK" description="" %}
```json
{
    "data": {
        "id": "string",
        "file_record_id": "string",
        "root_hash": "string",
        "cid": "string",
        "user_id": "string",
        "date_pinned": "2023-01-01T11:11:11.111111Z",
        "date_unpinned": "2023-11-11T11:11:11.111111Z",
        "pinned": false,
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
curl --location --request POST 'https://api-ipfs.attoaioz.cyou/api/pinning/pinByHash' \
--header 'pinning_api_key: KEY' \
--header 'pinning_secret_key: SECRET' \
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

url = "https://api-ipfs.attoaioz.cyou/api/pinning/pinByHash"

payload = json.dumps({
  "hash_to_pin": "Qmc1135ziMvmFG534i75E8HpJoLqzLzgKBxfjBV9cBsMAs",
  "metadata": {
    "name": "name-ipfs-hash"
  }
})
headers = {
  'pinning_api_key': 'KEY',
  'pinning_secret_key': 'SECRET',
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
