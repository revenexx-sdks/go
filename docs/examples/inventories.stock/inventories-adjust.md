```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/inventories_stock"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := inventories_stock.New(client)

response, error := service.InventoriesAdjust(
    inventories_stock.WithInventoriesAdjustItems([]interface{}{}),
    inventories_stock.WithInventoriesAdjustLocationCode("main"),
    inventories_stock.WithInventoriesAdjustProductId(""),
    inventories_stock.WithInventoriesAdjustQuantity(-3),
    inventories_stock.WithInventoriesAdjustReason("Stocktake 2026-03, two units damaged"),
    inventories_stock.WithInventoriesAdjustSku("ACME-4711-BLK"),
)
```
