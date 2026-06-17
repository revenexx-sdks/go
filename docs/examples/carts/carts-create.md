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

response, error := service.CartsCreate(
    carts.WithCartsCreateChannelId(""),
    carts.WithCartsCreateContactId(""),
    carts.WithCartsCreateCurrency(""),
    carts.WithCartsCreateIsCurrent(false),
    carts.WithCartsCreateMarketId(""),
    carts.WithCartsCreateMetadata(map[string]interface{}{}),
    carts.WithCartsCreateName(""),
    carts.WithCartsCreateSessionKey(""),
)
```
