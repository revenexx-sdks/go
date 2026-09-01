```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/inventories_locations"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := inventories_locations.New(client)

response, error := service.InventoriesLocationsList(
    inventories_locations.WithInventoriesLocationsListLimit(50),
    inventories_locations.WithInventoriesLocationsListOffset(0),
    inventories_locations.WithInventoriesLocationsListOrder("created_at.desc"),
    inventories_locations.WithInventoriesLocationsListId(""),
    inventories_locations.WithInventoriesLocationsListCode("main"),
    inventories_locations.WithInventoriesLocationsListName("Main warehouse"),
    inventories_locations.WithInventoriesLocationsListLabels("{}"),
    inventories_locations.WithInventoriesLocationsListType("warehouse"),
    inventories_locations.WithInventoriesLocationsListPriority(0),
    inventories_locations.WithInventoriesLocationsListEnabled(true),
    inventories_locations.WithInventoriesLocationsListAddress("{}"),
    inventories_locations.WithInventoriesLocationsListMetadata("{}"),
    inventories_locations.WithInventoriesLocationsListCreatedAt("2026-01-01T12:00:00Z"),
    inventories_locations.WithInventoriesLocationsListUpdatedAt("2026-01-01T12:00:00Z"),
)
```
