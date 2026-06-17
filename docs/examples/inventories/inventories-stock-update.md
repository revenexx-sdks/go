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

response, error := service.InventoriesStockUpdate(
    "",
    inventories.WithInventoriesStockUpdateLocationId(""),
    inventories.WithInventoriesStockUpdateMetadata(map[string]interface{}{}),
    inventories.WithInventoriesStockUpdateOnHand(0),
    inventories.WithInventoriesStockUpdateProductId(""),
    inventories.WithInventoriesStockUpdateReorderPoint(0),
    inventories.WithInventoriesStockUpdateReserved(0),
    inventories.WithInventoriesStockUpdateSku(""),
)
```
