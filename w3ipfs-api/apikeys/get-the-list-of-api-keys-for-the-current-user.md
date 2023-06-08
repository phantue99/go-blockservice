---
description: /api/apiKeys/list
---

# Get the list of API keys for the current user

{% swagger method="get" path="" baseUrl="https://api-ipfs.attoaioz.cyou/api/apiKeys/list" summary="" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" required="true" %}
JWT
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="" %}
```json
{
    "data": {
        "total": 1,
        "api_keys": [
            {
                "id": "string",
                "name": "string",
                "api_key": "string",
                "secret_key": "string",
                "scopes": {
                    "admin": true
                },
                "created_at": "2023-01-01T11:11:11.111111Z"
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
curl --location --request GET 'https://api-ipfs.attoaioz.cyou/api/apiKeys/list' \
--header 'Authorization: Bearer JWT' \
--data-raw ''
```
{% endtab %}

{% tab title="Node.js" %}
```javascript
var axios = require('axios');
var data = '';

var config = {
  method: 'get',
  url: 'https://api-ipfs.attoaioz.cyou/api/apiKeys/list',
  headers: { 
    'Authorization': 'Bearer JWT'
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

url = "https://api-ipfs.attoaioz.cyou/api/apiKeys/list"

payload = ""
headers = {
  'Authorization': 'Bearer JWT'
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

  url := "https://api-ipfs.attoaioz.cyou/api/apiKeys/list"
  method := "GET"

  payload := strings.NewReader(``)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Authorization", "Bearer JWT")

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
