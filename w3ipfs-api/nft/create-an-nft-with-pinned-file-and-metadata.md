---
description: /api/nft/
---

# Create an NFT with pinned file and metadata

{% swagger method="post" path="" baseUrl="https://api-ipfs.attoaioz.cyou/api/nft/" summary="" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="header" name="pinning_api_key" required="true" %}

{% endswagger-parameter %}

{% swagger-parameter in="header" name="pinning_secret_key" required="true" %}

{% endswagger-parameter %}

{% swagger-parameter in="body" required="true" name="file" type="File" %}

{% endswagger-parameter %}

{% swagger-parameter in="body" name="metadata" type="File" %}

{% endswagger-parameter %}

{% swagger-parameter in="query" name="cid" %}
asset cid
{% endswagger-parameter %}

{% swagger-response status="201: Created" description="" %}
```
{
    "data": {
        "id": "f0cda369-2846-414d-8e8d-8c0394113bef",
        "asset_cid": "bafkreiaopgqg3im4zzm2rq7mudz2qr3m62qrziatq3x73nlnidcmh3twjy",
        "metadata_cid": "bafkreigjjz4vizrrbvz6lu6fow3bkgnuuucrex46yuew74wq4s2r7mjaay",
        "asset_pin_id": "35e614f2-3161-4a4a-bb67-dbd68e2cc5ba",
        "metadata_pin_id": "71107573-58bc-427e-b536-05465e830dff",
        "size": 13557,
        "user_id": "c2452942-5dc5-4a0d-9c38-77a66edcf43a",
        "created_at": "2023-06-07T04:36:10.723127485Z",
        "updated_at": "2023-06-07T04:36:10.723127485Z",
        "pinned": true,
        "metadata_asset": {
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
curl --location --request POST 'https://api-ipfs.attoaioz.cyou/api/nft/' \
--header 'pinning_api_key: OOC4HF2dGRlgAVzg6vbypg==' \
--header 'pinning_secret_key: 0NFNueE1IKn0bbIMB8cRzG2/JeuIwc0BX/2exij8wco=' \
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
    'pinning_api_key': 'OOC4HF2dGRlgAVzg6vbypg==', 
    'pinning_secret_key': '0NFNueE1IKn0bbIMB8cRzG2/JeuIwc0BX/2exij8wco=', 
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
  'pinning_api_key': 'OOC4HF2dGRlgAVzg6vbypg==',
  'pinning_secret_key': '0NFNueE1IKn0bbIMB8cRzG2/JeuIwc0BX/2exij8wco='
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
  req.Header.Add("pinning_api_key", "OOC4HF2dGRlgAVzg6vbypg==")
  req.Header.Add("pinning_secret_key", "0NFNueE1IKn0bbIMB8cRzG2/JeuIwc0BX/2exij8wco=")

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
