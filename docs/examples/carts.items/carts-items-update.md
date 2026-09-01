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

response, error := service.CartsItemsUpdate(
    "",
    "",
    carts_items.WithCartsItemsUpdateConfiguration(map[string]interface{}{
        "colour": "RAL 7016",
        "finish": "brushed",
        "length_mm": 2400,
        "mounting": "wall"
    }),
    carts_items.WithCartsItemsUpdateCurrency("EUR"),
    carts_items.WithCartsItemsUpdateMetadata(map[string]interface{}{
        "campaign": "spring-catalogue",
        "locale": "de-DE",
        "source": "storefront"
    }),
    carts_items.WithCartsItemsUpdateName("Hex bolt M8"),
    carts_items.WithCartsItemsUpdatePosition(1),
    carts_items.WithCartsItemsUpdateProductId(""),
    carts_items.WithCartsItemsUpdateQuantity(9.99),
    carts_items.WithCartsItemsUpdateSku("BOLT-M8-30"),
    carts_items.WithCartsItemsUpdateSnapshot(map[string]interface{}{}),
    carts_items.WithCartsItemsUpdateTaxRate(19),
    carts_items.WithCartsItemsUpdateType("product"),
    carts_items.WithCartsItemsUpdateUnit("pcs"),
    carts_items.WithCartsItemsUpdateUnitPrice(9.99),
)
```
