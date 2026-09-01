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

response, error := service.ShippingTiersUpdate(
    "",
    "",
    shipping_methods.WithShippingTiersUpdateFromValue(10),
    shipping_methods.WithShippingTiersUpdatePosition(1),
    shipping_methods.WithShippingTiersUpdatePrice(6.9),
)
```
