# Users

## Get User Information

This API allows you to retrieve the details of the currently authenticated user.

{% swagger method="get" path="" baseUrl="https://api-ipfs.attoaioz.cyou/api/users/me" summary="" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" required="true" %}
JWT
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="The API will respond with a JSON object containing the user details." %}
```json
{
    "data": {
        "user": {
            "id": "string",
            "first_name": "string",
            "last_name": "string",
            "email": "string",
            "wallet_address": "string",
            "role": "string",
            "created_at": "2023-05-11T11:11:11.111111Z",
            "updated_at": "2023-06-11T11:11:11.111111Z"
        }
    },
    "status": "success"
}
```
{% endswagger-response %}
{% endswagger %}

{% tabs %}
{% tab title="cURL" %}
```
curl --location --request GET 'https://api-ipfs.attoaioz.cyou/api/users/me' \
--header 'Authorization: Bearer JWT'
```
{% endtab %}

{% tab title="Node.js" %}
```javascript
var axios = require('axios');

var config = {
  method: 'get',
  url: 'https://api-ipfs.attoaioz.cyou/api/users/me',
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

url = "https://api-ipfs.attoaioz.cyou/api/users/me"

payload={}
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
  "net/http"
  "io/ioutil"
)

func main() {

  url := "https://api-ipfs.attoaioz.cyou/api/users/me"
  method := "GET"

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

## Edit User

Edit the first name and last name of the user's profile

The request body when edit will look like this:&#x20;

```
{
    first_name: string,
    last_name: string
}
```

{% swagger method="put" path="" baseUrl="https://api-ipfs.attoaioz.cyou/api/users/editProfile" summary="" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" required="true" %}
JWT
{% endswagger-parameter %}

{% swagger-parameter in="body" name="first_name" %}
First name
{% endswagger-parameter %}

{% swagger-parameter in="body" name="last_name" %}
Last name
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="" %}
```json
{
    "message": "Profile updated successfully",
    "status": "success"
}
```
{% endswagger-response %}
{% endswagger %}

{% tabs %}
{% tab title="cURL" %}
```
curl --location --request PUT 'https://api-ipfs.attoaioz.cyou/api/users/editProfile' \
--header 'Authorization: Bearer JWT' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "First",
    "last_name": "Last"
}'
```
{% endtab %}

{% tab title="Node.js" %}
```javascript
var axios = require('axios');
var data = JSON.stringify({
  "first_name": "First",
  "last_name": "Last"
});

var config = {
  method: 'put',
  url: 'https://api-ipfs.attoaioz.cyou/api/users/editProfile',
  headers: { 
    'Authorization': 'Bearer JWT', 
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

url = "https://api-ipfs.attoaioz.cyou/api/users/editProfile"

payload = json.dumps({
  "first_name": "First",
  "last_name": "Last"
})
headers = {
  'Authorization': 'Bearer JWT',
  'Content-Type': 'application/json'
}

response = requests.request("PUT", url, headers=headers, data=payload)

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

  url := "https://api-ipfs.attoaioz.cyou/api/users/editProfile"
  method := "PUT"

  payload := strings.NewReader(`{
    "first_name": "First",
    "last_name": "Last"
}`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Authorization", "Bearer JWT")
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
