---
description: /api/nft/
---

# Create an NFT with pinned file and metadata

The metadata JSON file will look like this:

```
{
    "name": "My Awesome NFT",
    "description": "This is an NFT that represents my creativity as a digital artist!",
    "properties": [
        {
            "trait_type": "Color",
            "value": "Red"
        },
        {
            "trait_type": "Rarity",
            "value": "Medium"
        }
    ]
}
```

{% swagger method="post" path="" baseUrl="https://api-ipfs.attoaioz.cyou/api/nft/" summary="" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="header" name="pinning_api_key" required="true" %}
PINNING-API-KEY
{% endswagger-parameter %}

{% swagger-parameter in="header" name="pinning_secret_key" required="true" %}
PINNING-SECRET-KEY
{% endswagger-parameter %}

{% swagger-parameter in="body" required="true" name="file" type="File" %}

{% endswagger-parameter %}

{% swagger-parameter in="body" name="metadata" type="File" %}

{% endswagger-parameter %}

{% swagger-parameter in="query" name="cid" %}
asset cid
{% endswagger-parameter %}

{% swagger-response status="201: Created" description="" %}
```json
{
    "data": {
        "id": "string",
        "asset_cid": "string",
        "metadata_cid": "string",
        "asset_pin_id": "string",
        "metadata_pin_id": "string",
        "size": number,
        "user_id": "string",
        "created_at": "2023-01-01T11:11:11.111111Z",
        "updated_at": "2023-11-11T11:11:11.111111Z",
        "pinned": true,
        "metadata_asset": {
            "name": "string",
            "type": "string"
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
curl --location --request POST 'https://api-ipfs.attoaioz.cyou/api/nft/' \
--header 'pinning_api_key: KEY' \
--header 'pinning_secret_key: SECRET' \
--form 'metadata=@"/sample.json"' \
--form 'file=@"/test.png"'
```
{% endtab %}

{% tab title="Node.js" %}
```javascript
var axios = require('axios');
var FormData = require('form-data');
var fs = require('fs');
var data = new FormData();
data.append('metadata', fs.createReadStream('/sample.json'));
data.append('file', fs.createReadStream('/test.png'));

var config = {
  method: 'post',
  url: 'https://api-ipfs.attoaioz.cyou/api/nft/',
  headers: { 
    'pinning_api_key': 'KEY', 
    'pinning_secret_key': 'SECRET', 
    ...data.getHeaders()
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

url = "https://api-ipfs.attoaioz.cyou/api/nft/"

payload={}
files=[
  ('metadata',('sample.json',open('/sample.json','rb'),'application/json')),
  ('file',('test.png',open('/test.png','rb'),'image/png'))
]
headers = {
  'pinning_api_key': 'KEY',
  'pinning_secret_key': 'SECRET'
}

response = requests.request("POST", url, headers=headers, data=payload, files=files)

print(response.text)
```
{% endtab %}

{% tab title="Go" %}
```go
package main

import (
  "fmt"
  "bytes"
  "mime/multipart"
  "os"
  "path/filepath"
  "io"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "https://api-ipfs.attoaioz.cyou/api/nft/"
  method := "POST"

  payload := &bytes.Buffer{}
  writer := multipart.NewWriter(payload)
  file, errFile1 := os.Open("/sample.json")
  defer file.Close()
  part1, errFile1 := writer.CreateFormFile("metadata",filepath.Base("/sample.json"))
  _, errFile1 = io.Copy(part1, file)
  if errFile1 != nil {
    fmt.Println(errFile1)
    return
  }
  file, errFile2 := os.Open("/test.png")
  defer file.Close()
  part2, errFile2 := writer.CreateFormFile("file",filepath.Base("/test.png"))
  _, errFile2 = io.Copy(part2, file)
  if errFile2 != nil {
    fmt.Println(errFile2)
    return
  }
  err := writer.Close()
  if err != nil {
    fmt.Println(err)
    return
  }

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("pinning_api_key", "KEY")
  req.Header.Add("pinning_secret_key", "SECRET")

  req.Header.Set("Content-Type", writer.FormDataContentType())
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
