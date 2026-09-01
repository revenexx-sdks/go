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

response, error := service.InventoriesMovementsList(
    inventories_stock.WithInventoriesMovementsListLimit(50),
    inventories_stock.WithInventoriesMovementsListOffset(0),
    inventories_stock.WithInventoriesMovementsListOrder("created_at.desc"),
    inventories_stock.WithInventoriesMovementsListId(""),
    inventories_stock.WithInventoriesMovementsListLocationId(""),
    inventories_stock.WithInventoriesMovementsListProductId(""),
    inventories_stock.WithInventoriesMovementsListSku("ACME-4711-BLK"),
    inventories_stock.WithInventoriesMovementsListType("inbound"),
    inventories_stock.WithInventoriesMovementsListQuantity(5),
    inventories_stock.WithInventoriesMovementsListOrderRef("SO-2026-000123"),
    inventories_stock.WithInventoriesMovementsListReason("Delivery note 4711"),
    inventories_stock.WithInventoriesMovementsListMetadata("{}"),
    inventories_stock.WithInventoriesMovementsListCreatedAt("2026-01-01T12:00:00Z"),
)
```
