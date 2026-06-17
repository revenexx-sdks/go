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

response, error := service.InventoriesStockCreate(
    "",
    inventories.WithInventoriesStockCreateMetadata(map[string]interface{}{}),
    inventories.WithInventoriesStockCreateOnHand(0),
    inventories.WithInventoriesStockCreateProductId(""),
    inventories.WithInventoriesStockCreateReorderPoint(0),
    inventories.WithInventoriesStockCreateReserved(0),
    inventories.WithInventoriesStockCreateSku(""),
)
```
