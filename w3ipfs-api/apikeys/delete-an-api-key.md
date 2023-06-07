---
description: /api/apiKeys/:ID
---

# Delete an API Key

{% swagger method="delete" path=":ID" baseUrl="https://api-ipfs.attoaioz.cyou/api/apiKeys/" summary="" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" required="true" %}
JWT
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="" %}
```json
{
    "message": "API key has been deleted",
    "status": "success"
}
```
{% endswagger-response %}
{% endswagger %}

{% tabs %}
{% tab title="cURL" %}
```
curl --location --request DELETE 'https://api-ipfs.attoaioz.cyou/api/apiKeys/{ID}
--header 'Authorization: Bearer JWT'
```
{% endtab %}

{% tab title="Node.js" %}
```javascript
var axios = require('axios');

var config = {
  method: 'delete',
  url: 'https://api-ipfs.attoaioz.cyou/api/apiKeys/38ad2971-6be5-4dc3-ba28-890b9a86b8e9',
  headers: { 
    'Authorization': 'Bearer JWT'
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

url = "https://api-ipfs.attoaioz.cyou/api/apiKeys/38ad2971-6be5-4dc3-ba28-890b9a86b8e9"

payload={}
headers = {
  'Authorization': 'Bearer JWT'
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

  url := "https://api-ipfs.attoaioz.cyou/api/apiKeys/38ad2971-6be5-4dc3-ba28-890b9a86b8e9"
  method := "DELETE"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

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
