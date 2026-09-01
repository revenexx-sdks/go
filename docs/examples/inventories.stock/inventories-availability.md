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

response, error := service.InventoriesAvailability(
    inventories_stock.WithInventoriesAvailabilityItems([]interface{}{}),
    inventories_stock.WithInventoriesAvailabilityLocationCode("main"),
    inventories_stock.WithInventoriesAvailabilityProductId(""),
    inventories_stock.WithInventoriesAvailabilityQuantity(1),
    inventories_stock.WithInventoriesAvailabilitySku("ACME-4711-BLK"),
)
```
