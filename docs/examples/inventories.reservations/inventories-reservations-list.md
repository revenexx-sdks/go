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

response, error := service.InventoriesReservationsList(
    inventories_reservations.WithInventoriesReservationsListLimit(50),
    inventories_reservations.WithInventoriesReservationsListOffset(0),
    inventories_reservations.WithInventoriesReservationsListOrder("created_at.desc"),
    inventories_reservations.WithInventoriesReservationsListId(""),
    inventories_reservations.WithInventoriesReservationsListLocationId(""),
    inventories_reservations.WithInventoriesReservationsListProductId(""),
    inventories_reservations.WithInventoriesReservationsListSku("ACME-4711-BLK"),
    inventories_reservations.WithInventoriesReservationsListQuantity(2),
    inventories_reservations.WithInventoriesReservationsListOrderRef("SO-2026-000123"),
    inventories_reservations.WithInventoriesReservationsListStatus("active"),
    inventories_reservations.WithInventoriesReservationsListExpiresAt("2026-01-01T12:00:00Z"),
    inventories_reservations.WithInventoriesReservationsListMetadata("{}"),
    inventories_reservations.WithInventoriesReservationsListCreatedAt("2026-01-01T12:00:00Z"),
    inventories_reservations.WithInventoriesReservationsListUpdatedAt("2026-01-01T12:00:00Z"),
)
```
