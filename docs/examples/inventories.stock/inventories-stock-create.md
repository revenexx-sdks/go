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

response, error := service.InventoriesStockCreate(
    "",
    inventories_stock.WithInventoriesStockCreateMetadata(map[string]interface{}{
        "backorder": true
    }),
    inventories_stock.WithInventoriesStockCreateProductId(""),
    inventories_stock.WithInventoriesStockCreateReorderPoint(10),
    inventories_stock.WithInventoriesStockCreateSku("ACME-4711-BLK"),
)
```
