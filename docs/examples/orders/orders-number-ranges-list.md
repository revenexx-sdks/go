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

response, error := service.OrdersNumberRangesList(
    orders.WithOrdersNumberRangesListId(""),
    orders.WithOrdersNumberRangesListCode("order"),
    orders.WithOrdersNumberRangesListPrefix("ORD-"),
    orders.WithOrdersNumberRangesListSuffix(""),
    orders.WithOrdersNumberRangesListPadding(6),
    orders.WithOrdersNumberRangesListCounter(123),
    orders.WithOrdersNumberRangesListStep(1),
    orders.WithOrdersNumberRangesListPositionStep(10),
    orders.WithOrdersNumberRangesListChannelId(""),
    orders.WithOrdersNumberRangesListCreatedAt("2026-01-01T12:00:00Z"),
    orders.WithOrdersNumberRangesListUpdatedAt("2026-01-01T12:00:00Z"),
    orders.WithOrdersNumberRangesListLimit(50),
    orders.WithOrdersNumberRangesListOffset(0),
    orders.WithOrdersNumberRangesListOrder("created_at.desc"),
)
```
