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

response, error := service.CartsIoProfilesList(
    carts_io.WithCartsIoProfilesListId(""),
    carts_io.WithCartsIoProfilesListName("cart-export-csv"),
    carts_io.WithCartsIoProfilesListDirection("import"),
    carts_io.WithCartsIoProfilesListEntity("carts"),
    carts_io.WithCartsIoProfilesListFormat("json"),
    carts_io.WithCartsIoProfilesListApplyMode("insert"),
    carts_io.WithCartsIoProfilesListIsTemplate(true),
    carts_io.WithCartsIoProfilesListCreatedAt("2026-01-01T12:00:00Z"),
    carts_io.WithCartsIoProfilesListUpdatedAt("2026-01-01T12:00:00Z"),
    carts_io.WithCartsIoProfilesListLimit(1),
    carts_io.WithCartsIoProfilesListOffset(1),
    carts_io.WithCartsIoProfilesListOrder("created_at.desc"),
)
```
