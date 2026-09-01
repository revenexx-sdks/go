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

response, error := service.ShippingTiersCreate(
    "",
    shipping_methods.WithShippingTiersCreateFromValue(10),
    shipping_methods.WithShippingTiersCreatePosition(1),
    shipping_methods.WithShippingTiersCreatePrice(6.9),
)
```
