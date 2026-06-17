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

response, error := service.CartsItemsUpdate(
    "",
    "",
    carts.WithCartsItemsUpdateConfiguration(map[string]interface{}{}),
    carts.WithCartsItemsUpdateCurrency(""),
    carts.WithCartsItemsUpdateMetadata(map[string]interface{}{}),
    carts.WithCartsItemsUpdateName(""),
    carts.WithCartsItemsUpdatePosition(0),
    carts.WithCartsItemsUpdateProductId(""),
    carts.WithCartsItemsUpdateQuantity(0),
    carts.WithCartsItemsUpdateSku(""),
    carts.WithCartsItemsUpdateSnapshot(map[string]interface{}{}),
    carts.WithCartsItemsUpdateTaxRate(0),
    carts.WithCartsItemsUpdateType(""),
    carts.WithCartsItemsUpdateUnit(""),
    carts.WithCartsItemsUpdateUnitPrice(0),
)
```
