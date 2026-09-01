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

response, error := service.InventoriesStockUpdate(
    "",
    inventories_stock.WithInventoriesStockUpdateLocationId(""),
    inventories_stock.WithInventoriesStockUpdateMetadata(map[string]interface{}{
        "backorder": true
    }),
    inventories_stock.WithInventoriesStockUpdateProductId(""),
    inventories_stock.WithInventoriesStockUpdateReorderPoint(10),
    inventories_stock.WithInventoriesStockUpdateSku("ACME-4711-BLK"),
)
```
