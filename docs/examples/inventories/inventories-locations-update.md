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

response, error := service.InventoriesLocationsUpdate(
    "",
    inventories.WithInventoriesLocationsUpdateAddress(map[string]interface{}{}),
    inventories.WithInventoriesLocationsUpdateCode(""),
    inventories.WithInventoriesLocationsUpdateEnabled(false),
    inventories.WithInventoriesLocationsUpdateLabels(map[string]interface{}{}),
    inventories.WithInventoriesLocationsUpdateMetadata(map[string]interface{}{}),
    inventories.WithInventoriesLocationsUpdateName(""),
    inventories.WithInventoriesLocationsUpdatePriority(0),
    inventories.WithInventoriesLocationsUpdateType(""),
)
```
