```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/shipping_value_lists"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := shipping_value_lists.New(client)

response, error := service.ShippingWeightUnitsList(
    shipping_value_lists.WithShippingWeightUnitsListLimit(1),
    shipping_value_lists.WithShippingWeightUnitsListOffset(1),
)
```
