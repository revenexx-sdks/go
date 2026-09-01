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

response, error := service.InventoriesLocationsUpdate(
    "",
    inventories_locations.WithInventoriesLocationsUpdateAddress(map[string]interface{}{
        "city": "Nuremberg",
        "country": "DE",
        "postal_code": "90402",
        "street": "Industriering 4"
    }),
    inventories_locations.WithInventoriesLocationsUpdateCode("main"),
    inventories_locations.WithInventoriesLocationsUpdateEnabled(true),
    inventories_locations.WithInventoriesLocationsUpdateLabels(map[string]interface{}{
        "de": "Hauptlager",
        "en": "Main warehouse"
    }),
    inventories_locations.WithInventoriesLocationsUpdateMetadata(map[string]interface{}{
        "erp_site": "1000"
    }),
    inventories_locations.WithInventoriesLocationsUpdateName("Main warehouse"),
    inventories_locations.WithInventoriesLocationsUpdatePriority(0),
    inventories_locations.WithInventoriesLocationsUpdateType("warehouse"),
)
```
