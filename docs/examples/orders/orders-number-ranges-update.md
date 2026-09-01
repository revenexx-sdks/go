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

response, error := service.OrdersNumberRangesUpdate(
    "",
    orders.WithOrdersNumberRangesUpdateChannelId(""),
    orders.WithOrdersNumberRangesUpdateCode("order"),
    orders.WithOrdersNumberRangesUpdateCounter(123),
    orders.WithOrdersNumberRangesUpdateMetadata(map[string]interface{}{
        "owner": "erp-sync"
    }),
    orders.WithOrdersNumberRangesUpdatePadding(6),
    orders.WithOrdersNumberRangesUpdatePositionStep(10),
    orders.WithOrdersNumberRangesUpdatePrefix("ORD-"),
    orders.WithOrdersNumberRangesUpdateStep(1),
    orders.WithOrdersNumberRangesUpdateSuffix(""),
)
```
