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

response, error := service.CartsCreate(
    carts.WithCartsCreateChannelId(""),
    carts.WithCartsCreateContactId(""),
    carts.WithCartsCreateCurrency("EUR"),
    carts.WithCartsCreateIsCurrent(true),
    carts.WithCartsCreateMetadata(map[string]interface{}{
        "campaign": "spring-catalogue",
        "locale": "de-DE",
        "source": "storefront"
    }),
    carts.WithCartsCreateName("Weekly order"),
    carts.WithCartsCreateSessionKey("a1b2c3d4e5f6"),
)
```
