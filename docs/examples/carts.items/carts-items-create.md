```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/carts_items"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := carts_items.New(client)

response, error := service.CartsItemsCreate(
    "",
    carts_items.WithCartsItemsCreateConfiguration(map[string]interface{}{
        "colour": "RAL 7016",
        "finish": "brushed",
        "length_mm": 2400,
        "mounting": "wall"
    }),
    carts_items.WithCartsItemsCreateCurrency("EUR"),
    carts_items.WithCartsItemsCreateMetadata(map[string]interface{}{
        "campaign": "spring-catalogue",
        "locale": "de-DE",
        "source": "storefront"
    }),
    carts_items.WithCartsItemsCreateName("Hex bolt M8"),
    carts_items.WithCartsItemsCreatePosition(1),
    carts_items.WithCartsItemsCreateProductId(""),
    carts_items.WithCartsItemsCreateQuantity(9.99),
    carts_items.WithCartsItemsCreateSku("BOLT-M8-30"),
    carts_items.WithCartsItemsCreateSnapshot(map[string]interface{}{}),
    carts_items.WithCartsItemsCreateTaxRate(19),
    carts_items.WithCartsItemsCreateType("product"),
    carts_items.WithCartsItemsCreateUnit("pcs"),
    carts_items.WithCartsItemsCreateUnitPrice(9.99),
)
```
