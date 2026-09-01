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

response, error := service.CartsIoProfilesUpdate(
    "",
    carts_io.WithCartsIoProfilesUpdateApplyMode("insert"),
    carts_io.WithCartsIoProfilesUpdateDirection("import"),
    carts_io.WithCartsIoProfilesUpdateEntity("carts"),
    carts_io.WithCartsIoProfilesUpdateFormat("json"),
    carts_io.WithCartsIoProfilesUpdateIsTemplate(true),
    carts_io.WithCartsIoProfilesUpdateMapping(map[string]interface{}{}),
    carts_io.WithCartsIoProfilesUpdateName("cart-export-csv"),
    carts_io.WithCartsIoProfilesUpdateOptions(map[string]interface{}{}),
)
```
