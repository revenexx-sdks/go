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

response, error := service.InventoriesRestock(
    inventories_stock.WithInventoriesRestockItems([]interface{}{}),
    inventories_stock.WithInventoriesRestockLocationCode("main"),
    inventories_stock.WithInventoriesRestockOrderRef("SO-2026-000123"),
    inventories_stock.WithInventoriesRestockProductId(""),
    inventories_stock.WithInventoriesRestockQuantity(1),
    inventories_stock.WithInventoriesRestockReason("Return: wrong size"),
    inventories_stock.WithInventoriesRestockRestock(true),
    inventories_stock.WithInventoriesRestockSku("ACME-4711-BLK"),
)
```
