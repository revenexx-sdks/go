```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/shipping_methods"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := shipping_methods.New(client)

response, error := service.ShippingTiersLadder(
    "",
    4.9,
    5,
    30,
    shipping_methods.WithShippingTiersLadderFromValue(0),
    shipping_methods.WithShippingTiersLadderReplace(true),
    shipping_methods.WithShippingTiersLadderStepPrice(2),
)
```
