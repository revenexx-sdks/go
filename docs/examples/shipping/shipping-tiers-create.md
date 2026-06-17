```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/shipping"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com")
    client.WithApiKeyAuth("<API_KEY>")
)

service := shipping.New(client)

response, error := service.ShippingTiersCreate(
    "",
    shipping.WithShippingTiersCreateFromValue(0),
    shipping.WithShippingTiersCreatePosition(0),
    shipping.WithShippingTiersCreatePrice(0),
)
```
