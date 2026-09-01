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

response, error := service.ShippingWeightUnitsUpdate(
    "",
    shipping_value_lists.WithShippingWeightUnitsUpdateDescription("When to pick this weight unit."),
    shipping_value_lists.WithShippingWeightUnitsUpdateDescriptions(map[string]interface{}{
        "de": "Wann diese Option zu w\u00e4hlen ist.",
        "en": "When to pick this weight unit."
    }),
    shipping_value_lists.WithShippingWeightUnitsUpdateFactor(1000),
    shipping_value_lists.WithShippingWeightUnitsUpdateIsDefault(true),
    shipping_value_lists.WithShippingWeightUnitsUpdateLabels(map[string]interface{}{
        "de": "Tonne",
        "en": "Tonne"
    }),
    shipping_value_lists.WithShippingWeightUnitsUpdatePosition(1),
    shipping_value_lists.WithShippingWeightUnitsUpdateTitle("Tonne"),
    shipping_value_lists.WithShippingWeightUnitsUpdateTone("neutral"),
)
```
