```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/orders"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := orders.New(client)

response, error := service.OrdersReturn(
    "",
    orders.WithOrdersReturnMetadata(map[string]interface{}{
        "rma_portal_case": "C-2026-0917"
    }),
    orders.WithOrdersReturnPositions([]interface{}{}),
    orders.WithOrdersReturnReason("Damaged on arrival"),
    orders.WithOrdersReturnRestock(true),
)
```
