```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/carts"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com")
    client.WithApiKeyAuth("<API_KEY>")
)

service := carts.New(client)

response, error := service.CartsIoProfilesCreate(
    "",
    "",
    carts.WithCartsIoProfilesCreateApplyMode(""),
    carts.WithCartsIoProfilesCreateEntity(""),
    carts.WithCartsIoProfilesCreateFormat(""),
    carts.WithCartsIoProfilesCreateIsTemplate(false),
    carts.WithCartsIoProfilesCreateMapping(map[string]interface{}{}),
    carts.WithCartsIoProfilesCreateOptions(map[string]interface{}{}),
)
```
