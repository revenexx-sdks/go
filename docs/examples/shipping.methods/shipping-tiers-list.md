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

response, error := service.ShippingTiersList(
    "",
    shipping_methods.WithShippingTiersListLimit(1),
    shipping_methods.WithShippingTiersListOffset(1),
    shipping_methods.WithShippingTiersListOrder("position.asc"),
    shipping_methods.WithShippingTiersListFromValue(10),
)
```
