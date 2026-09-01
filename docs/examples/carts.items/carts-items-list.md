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

response, error := service.CartsItemsList(
    "",
    carts_items.WithCartsItemsListId(""),
    carts_items.WithCartsItemsListType("product"),
    carts_items.WithCartsItemsListProductId(""),
    carts_items.WithCartsItemsListSku("BOLT-M8-30"),
    carts_items.WithCartsItemsListName("Hex bolt M8"),
    carts_items.WithCartsItemsListQuantity(100),
    carts_items.WithCartsItemsListUnit("pcs"),
    carts_items.WithCartsItemsListUnitPrice(0.12),
    carts_items.WithCartsItemsListCurrency("EUR"),
    carts_items.WithCartsItemsListTaxRate(19),
    carts_items.WithCartsItemsListLineTotal(12),
    carts_items.WithCartsItemsListPosition(0),
    carts_items.WithCartsItemsListCreatedAt("2026-01-01T12:00:00Z"),
    carts_items.WithCartsItemsListUpdatedAt("2026-01-01T12:00:00Z"),
    carts_items.WithCartsItemsListLimit(1),
    carts_items.WithCartsItemsListOffset(1),
    carts_items.WithCartsItemsListOrder("created_at.desc"),
)
```
