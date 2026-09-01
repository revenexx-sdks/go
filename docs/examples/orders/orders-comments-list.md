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

response, error := service.OrdersCommentsList(
    "",
    orders.WithOrdersCommentsListIdQuery(""),
    orders.WithOrdersCommentsListBody("Called the customer, delivery agreed for next week."),
    orders.WithOrdersCommentsListVisibility("internal"),
    orders.WithOrdersCommentsListAuthor("service-desk"),
    orders.WithOrdersCommentsListCreatedAt("2026-01-01T12:00:00Z"),
    orders.WithOrdersCommentsListLimit(50),
    orders.WithOrdersCommentsListOffset(0),
    orders.WithOrdersCommentsListOrder("created_at.desc"),
)
```
