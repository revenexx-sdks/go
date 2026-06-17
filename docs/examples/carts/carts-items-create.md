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

response, error := service.CartsItemsCreate(
    "",
    carts.WithCartsItemsCreateConfiguration(map[string]interface{}{}),
    carts.WithCartsItemsCreateCurrency(""),
    carts.WithCartsItemsCreateMetadata(map[string]interface{}{}),
    carts.WithCartsItemsCreateName(""),
    carts.WithCartsItemsCreatePosition(0),
    carts.WithCartsItemsCreateProductId(""),
    carts.WithCartsItemsCreateQuantity(0),
    carts.WithCartsItemsCreateSku(""),
    carts.WithCartsItemsCreateSnapshot(map[string]interface{}{}),
    carts.WithCartsItemsCreateTaxRate(0),
    carts.WithCartsItemsCreateType(""),
    carts.WithCartsItemsCreateUnit(""),
    carts.WithCartsItemsCreateUnitPrice(0),
)
```
