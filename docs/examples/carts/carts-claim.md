```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/carts"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := carts.New(client)

response, error := service.CartsClaim(
    "",
    "a1b2c3d4e5f6",
    carts.WithCartsClaimStrategy("merge"),
    carts.WithCartsClaimTargetCartId(""),
)
```
