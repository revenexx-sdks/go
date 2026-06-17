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

response, error := service.CartsIoProfilesUpdate(
    "",
    carts.WithCartsIoProfilesUpdateApplyMode(""),
    carts.WithCartsIoProfilesUpdateDirection(""),
    carts.WithCartsIoProfilesUpdateEntity(""),
    carts.WithCartsIoProfilesUpdateFormat(""),
    carts.WithCartsIoProfilesUpdateIsTemplate(false),
    carts.WithCartsIoProfilesUpdateMapping(map[string]interface{}{}),
    carts.WithCartsIoProfilesUpdateName(""),
    carts.WithCartsIoProfilesUpdateOptions(map[string]interface{}{}),
)
```
