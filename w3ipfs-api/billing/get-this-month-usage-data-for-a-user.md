---
description: /api/billing/thisMonthUsage/
---

# Get this month usage data for a user

{% swagger method="get" path="" baseUrl="https://api-ipfs.attoaioz.cyou/api/billing/thisMonthUsage" summary="" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="query" name="offset" %}
(default: 0)
{% endswagger-parameter %}

{% swagger-parameter in="header" name="pinning_api_key" required="true" %}
PINNING-API-KEY
{% endswagger-parameter %}

{% swagger-parameter in="query" name="limit" %}
(default: 10)
{% endswagger-parameter %}

{% swagger-parameter in="header" name="pinning_secret_key" required="true" %}
PINNING-SECRET-KEY
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="" %}
```json
{
    "data": {
        "total_storage": number,
        "storage_usage": number,
        "bandwidth_usage": number,
        "storage_cost": "string",
        "bandwidth_cost": "string"
    },
    "status": "success"
}
```
{% endswagger-response %}
{% endswagger %}

{% tabs %}
{% tab title="cURL" %}
```
curl --location --request GET 'https://api-ipfs.attoaioz.cyou/api/billing/thisMonthUsage?offset=0&limit=10' \
--header 'pinning_api_key: KEY' \
--header 'pinning_secret_key: SECRET'
```
{% endtab %}

{% tab title="Node.js" %}
```javascript
var axios = require('axios');

var config = {
  method: 'get',
  url: 'https://api-ipfs.attoaioz.cyou/api/billing/thisMonthUsage?offset=0&limit=10',
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

url = "https://api-ipfs.attoaioz.cyou/api/billing/thisMonthUsage?offset=0&limit=10"

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

  url := "https://api-ipfs.attoaioz.cyou/api/billing/thisMonthUsage?offset=0&limit=10"
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
