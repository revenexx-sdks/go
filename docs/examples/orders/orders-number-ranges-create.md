```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/orders"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com")
    client.WithApiKeyAuth("<API_KEY>")
)

service := orders.New(client)

response, error := service.OrdersNumberRangesCreate(
    "",
    orders.WithOrdersNumberRangesCreateChannelId(""),
    orders.WithOrdersNumberRangesCreateCounter(0),
    orders.WithOrdersNumberRangesCreateMetadata(map[string]interface{}{}),
    orders.WithOrdersNumberRangesCreatePadding(0),
    orders.WithOrdersNumberRangesCreatePositionStep(0),
    orders.WithOrdersNumberRangesCreatePrefix(""),
    orders.WithOrdersNumberRangesCreateStep(0),
    orders.WithOrdersNumberRangesCreateSuffix(""),
)
```
