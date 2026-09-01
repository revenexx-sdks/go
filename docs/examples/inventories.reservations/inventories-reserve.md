```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/inventories_reservations"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := inventories_reservations.New(client)

response, error := service.InventoriesReserve(
    "SO-2026-000123",
    inventories_reservations.WithInventoriesReserveExpiresAt("2026-01-01T12:00:00Z"),
    inventories_reservations.WithInventoriesReserveItems([]interface{}{}),
    inventories_reservations.WithInventoriesReserveLocationCode("main"),
    inventories_reservations.WithInventoriesReserveProductId(""),
    inventories_reservations.WithInventoriesReserveQuantity(2),
    inventories_reservations.WithInventoriesReserveShipTo(map[string]interface{}{}),
    inventories_reservations.WithInventoriesReserveSku("ACME-4711-BLK"),
)
```
