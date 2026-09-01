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

response, error := service.InventoriesStockList(
    inventories_stock.WithInventoriesStockListLimit(50),
    inventories_stock.WithInventoriesStockListOffset(0),
    inventories_stock.WithInventoriesStockListOrder("created_at.desc"),
    inventories_stock.WithInventoriesStockListId(""),
    inventories_stock.WithInventoriesStockListLocationId(""),
    inventories_stock.WithInventoriesStockListProductId(""),
    inventories_stock.WithInventoriesStockListSku("ACME-4711-BLK"),
    inventories_stock.WithInventoriesStockListOnHand(42),
    inventories_stock.WithInventoriesStockListReserved(5),
    inventories_stock.WithInventoriesStockListReorderPoint(10),
    inventories_stock.WithInventoriesStockListMetadata("{}"),
    inventories_stock.WithInventoriesStockListCreatedAt("2026-01-01T12:00:00Z"),
    inventories_stock.WithInventoriesStockListUpdatedAt("2026-01-01T12:00:00Z"),
)
```
