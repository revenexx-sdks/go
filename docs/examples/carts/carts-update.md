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

response, error := service.CartsUpdate(
    "",
    carts.WithCartsUpdateChannelId(""),
    carts.WithCartsUpdateCurrency("EUR"),
    carts.WithCartsUpdateMetadata(map[string]interface{}{
        "campaign": "spring-catalogue",
        "locale": "de-DE",
        "source": "storefront"
    }),
    carts.WithCartsUpdateName("Weekly order"),
)
```
