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

response, error := service.InventoriesReceive(
    []interface{}{},
    inventories.WithInventoriesReceiveLocationCode(""),
    inventories.WithInventoriesReceiveReason(""),
)
```
