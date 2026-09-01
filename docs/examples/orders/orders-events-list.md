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

response, error := service.OrdersEventsList(
    "",
    orders.WithOrdersEventsListIdQuery(""),
    orders.WithOrdersEventsListName("order.shipment.created"),
    orders.WithOrdersEventsListActor(""),
    orders.WithOrdersEventsListCreatedAt("2026-01-01T12:00:00Z"),
    orders.WithOrdersEventsListLimit(50),
    orders.WithOrdersEventsListOffset(0),
    orders.WithOrdersEventsListOrder("created_at.desc"),
)
```
