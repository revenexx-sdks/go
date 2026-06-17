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

response, error := service.OrdersNumberRangesUpdate(
    "",
    orders.WithOrdersNumberRangesUpdateChannelId(""),
    orders.WithOrdersNumberRangesUpdateCode(""),
    orders.WithOrdersNumberRangesUpdateCounter(0),
    orders.WithOrdersNumberRangesUpdateMetadata(map[string]interface{}{}),
    orders.WithOrdersNumberRangesUpdatePadding(0),
    orders.WithOrdersNumberRangesUpdatePositionStep(0),
    orders.WithOrdersNumberRangesUpdatePrefix(""),
    orders.WithOrdersNumberRangesUpdateStep(0),
    orders.WithOrdersNumberRangesUpdateSuffix(""),
)
```
