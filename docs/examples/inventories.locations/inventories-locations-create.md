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

response, error := service.InventoriesLocationsCreate(
    "main",
    "Main warehouse",
    inventories_locations.WithInventoriesLocationsCreateAddress(map[string]interface{}{
        "city": "Nuremberg",
        "country": "DE",
        "postal_code": "90402",
        "street": "Industriering 4"
    }),
    inventories_locations.WithInventoriesLocationsCreateEnabled(true),
    inventories_locations.WithInventoriesLocationsCreateLabels(map[string]interface{}{
        "de": "Hauptlager",
        "en": "Main warehouse"
    }),
    inventories_locations.WithInventoriesLocationsCreateMetadata(map[string]interface{}{
        "erp_site": "1000"
    }),
    inventories_locations.WithInventoriesLocationsCreatePriority(0),
    inventories_locations.WithInventoriesLocationsCreateType("warehouse"),
)
```
