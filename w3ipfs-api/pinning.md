# Pinning

## Pin a file or folder to IPFS

This API allows you to pin a file to IPFS using the provided pinning API key and secret key.

Example metadata:

```
{"name": "sample name", "keyvalues":{"key1": "value1","key2": "value2"}}
```

{% swagger method="post" path="" baseUrl="https://api-ipfs.attoaioz.cyou/api/pinning" summary="" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="header" name="pinning_api_key" required="true" %}
PINNING-API-KEY
{% endswagger-parameter %}

{% swagger-parameter in="body" name="metadata" %}
Optional stringified object
{% endswagger-parameter %}

{% swagger-parameter in="header" name="pinning_secret_key" required="true" %}
PINNING-SECRET-KEY
{% endswagger-parameter %}

{% swagger-parameter in="body" name="file" required="true" %}
Read stream representing the file
{% endswagger-parameter %}

{% swagger-response status="201: Created" description="" %}
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
        "is_pin_by_hash": false,
        "sub_hash_status": "string",
        "is_dir": false,
        "metadata": {
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
curl --location --request POST 'https://api-ipfs.attoaioz.cyou/api/pinning/' \
--header 'pinning_api_key: KEY' \
--header 'pinning_secret_key: SECRET' \
--form 'file=@"/test.png"'
```
{% endtab %}

{% tab title="Node.js" %}
```javascript
var axios = require('axios');
var FormData = require('form-data');
var fs = require('fs');
var data = new FormData();
data.append('file', fs.createReadStream('/test.png'));

var config = {
  method: 'post',
  url: 'https://api-ipfs.attoaioz.cyou/api/pinning/',
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

url = "https://api-ipfs.attoaioz.cyou/api/pinning/"

payload={}
files=[
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

  url := "https://api-ipfs.attoaioz.cyou/api/pinning/"
  method := "POST"

  payload := &bytes.Buffer{}
  writer := multipart.NewWriter(payload)
  file, errFile1 := os.Open("/test.png")
  defer file.Close()
  part1,
         errFile1 := writer.CreateFormFile("file",filepath.Base("/test.png"))
  _, errFile1 = io.Copy(part1, file)
  if errFile1 != nil {
    fmt.Println(errFile1)
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

## Pin a file to IPFS by its CID

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

## Retrieve a pin by its ID

{% swagger method="get" path="" baseUrl="https://api-ipfs.attoaioz.cyou/api/pinning/:pinId" summary="" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="header" name="pinning_api_key" required="true" %}
PINNING-API-KEY
{% endswagger-parameter %}

{% swagger-parameter in="header" name="pinning_secret_key" required="true" %}
PINNING-SECRET-KEY
{% endswagger-parameter %}

{% swagger-parameter in="path" required="true" name="pinId" %}
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
        "is_pin_by_hash": false,
        "sub_hash_status": "string",
        "is_dir": false,
        "metadata": {
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
curl --location --request GET 'https://api-ipfs.attoaioz.cyou/api/pinning/3c3fec2a-ca65-4b8e-bcf3-c8e2ceaa23d2' \
--header 'pinning_api_key: KEY' \
--header 'pinning_secret_key: SECRET' 
```
{% endtab %}

{% tab title="Node.js" %}
```javascript
var axios = require('axios');

var config = {
  method: 'get',
  url: 'https://api-ipfs.attoaioz.cyou/api/pinning/3c3fec2a-ca65-4b8e-bcf3-c8e2ceaa23d2',
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

url = "https://api-ipfs.attoaioz.cyou/api/pinning/3c3fec2a-ca65-4b8e-bcf3-c8e2ceaa23d2"

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

  url := "https://api-ipfs.attoaioz.cyou/api/pinning/3c3fec2a-ca65-4b8e-bcf3-c8e2ceaa23d2"
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

## Retrieve pins associated with a user



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
--header 'pinning_api_key: KEY' \
--header 'pinning_secret_key: SECRET'
```
{% endtab %}

{% tab title="Node.js" %}
```javascript
var axios = require('axios');

var config = {
  method: 'get',
  url: 'https://api-ipfs.attoaioz.cyou/api/pinning/pins/?offset=0&limit=10&pinned=true&sortBy=name&sortOrder=ASC',
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

url = "https://api-ipfs.attoaioz.cyou/api/pinning/pins/?offset=0&limit=10&pinned=true&sortBy=name&sortOrder=ASC"

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

  url := "https://api-ipfs.attoaioz.cyou/api/pinning/pins/?offset=0&limit=10&pinned=true&sortBy=name&sortOrder=ASC"
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

## Unpin a file from IPFS

{% swagger method="delete" path="" baseUrl="https://api-ipfs.attoaioz.cyou/api/pinning/unpin/:pinId" summary="" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="header" name="pinning_api_key" %}
PINNING-API-KEY
{% endswagger-parameter %}

{% swagger-parameter in="header" name="pinning_secret_key" %}
PINNING-SECRET-KEY
{% endswagger-parameter %}

{% swagger-parameter in="path" required="true" name="pinId" %}
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
