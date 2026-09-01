```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/carts_io"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := carts_io.New(client)

response, error := service.CartsIoProfilesCreate(
    "import",
    "cart-export-csv",
    carts_io.WithCartsIoProfilesCreateApplyMode("insert"),
    carts_io.WithCartsIoProfilesCreateEntity("carts"),
    carts_io.WithCartsIoProfilesCreateFormat("json"),
    carts_io.WithCartsIoProfilesCreateIsTemplate(true),
    carts_io.WithCartsIoProfilesCreateMapping(map[string]interface{}{}),
    carts_io.WithCartsIoProfilesCreateOptions(map[string]interface{}{}),
)
```
