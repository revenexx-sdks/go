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

response, error := service.OrdersShip(
    "",
    orders.WithOrdersShipCarrier("DHL"),
    orders.WithOrdersShipMetadata(map[string]interface{}{
        "warehouse": "HAM-1"
    }),
    orders.WithOrdersShipNumber("DEL-000123"),
    orders.WithOrdersShipPositions([]interface{}{}),
    orders.WithOrdersShipShippedAt("2026-01-01T12:00:00Z"),
    orders.WithOrdersShipTrackingCode("00340434161234567890"),
    orders.WithOrdersShipTrackingUrl("https://example.com/track/00340434161234567890"),
)
```
