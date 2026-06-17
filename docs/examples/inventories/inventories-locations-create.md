```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/inventories"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com")
    client.WithApiKeyAuth("<API_KEY>")
)

service := inventories.New(client)

response, error := service.InventoriesLocationsCreate(
    "",
    "",
    inventories.WithInventoriesLocationsCreateAddress(map[string]interface{}{}),
    inventories.WithInventoriesLocationsCreateEnabled(false),
    inventories.WithInventoriesLocationsCreateLabels(map[string]interface{}{}),
    inventories.WithInventoriesLocationsCreateMetadata(map[string]interface{}{}),
    inventories.WithInventoriesLocationsCreatePriority(0),
    inventories.WithInventoriesLocationsCreateType(""),
)
```
