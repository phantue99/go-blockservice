# Gateway

## Get list public & dedicated gateways

The IPFS API provides a list of public and dedicated gateways that you can use to access content on the IPFS network. Here is how you can retrieve the list of gateways using the provided API endpoint:



{% swagger method="get" path="" baseUrl="https://api-ipfs.attoaioz.cyou/api/gateways/" summary="" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="query" name="offset" %}
(default 0)
{% endswagger-parameter %}

{% swagger-parameter in="query" name="limit" %}
(default 10)
{% endswagger-parameter %}

{% swagger-parameter in="header" name="pinning_api_key" required="true" %}
PINNING-API-KEY
{% endswagger-parameter %}

{% swagger-parameter in="query" name="type" %}
Filter by type (options: Public, Dedicated, all) (default all)
{% endswagger-parameter %}

{% swagger-parameter in="header" name="pinning_secret_key" required="true" %}
PINNING-SECRET-KEY
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="The response will be in JSON format with the following structure:" %}
```json
{
    "data": {
        "totals": number,
        "gateways": [
            {
                "name": "string",
                "host": "string",
                "type": "string",
                "bandwidth": number,
                "operation": "string",
                "active": true
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
curl --location --request GET 'https://api-ipfs.attoaioz.cyou/api/gateways/?offset=0&limit=10&type=all' \
--header 'pinning_api_key: KEY' \
--header 'pinning_secret_key: SECRET'
```
{% endtab %}

{% tab title="Node.js" %}
```javascript
var axios = require('axios');

var config = {
  method: 'get',
  url: 'https://api-ipfs.attoaioz.cyou/api/gateways/?offset=0&limit=10&type=all',
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

url = "https://api-ipfs.attoaioz.cyou/api/gateways/?offset=0&limit=10&type=all"

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

  url := "https://api-ipfs.attoaioz.cyou/api/gateways/?offset=0&limit=10&type=all"
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

Please replace `YOUR_API_KEY` and `YOUR_SECRET_KEY` with your actual pinning API key and secret key.
