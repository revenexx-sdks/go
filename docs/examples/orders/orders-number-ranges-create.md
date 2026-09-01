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

response, error := service.OrdersNumberRangesCreate(
    "order",
    orders.WithOrdersNumberRangesCreateChannelId(""),
    orders.WithOrdersNumberRangesCreateCounter(123),
    orders.WithOrdersNumberRangesCreateMetadata(map[string]interface{}{
        "owner": "erp-sync"
    }),
    orders.WithOrdersNumberRangesCreatePadding(6),
    orders.WithOrdersNumberRangesCreatePositionStep(10),
    orders.WithOrdersNumberRangesCreatePrefix("ORD-"),
    orders.WithOrdersNumberRangesCreateStep(1),
    orders.WithOrdersNumberRangesCreateSuffix(""),
)
```
