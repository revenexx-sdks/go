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

response, error := service.ShippingWeightUnitsCreate(
    "t",
    1000,
    "Tonne",
    shipping_value_lists.WithShippingWeightUnitsCreateDescription("When to pick this weight unit."),
    shipping_value_lists.WithShippingWeightUnitsCreateDescriptions(map[string]interface{}{
        "de": "Wann diese Option zu w\u00e4hlen ist.",
        "en": "When to pick this weight unit."
    }),
    shipping_value_lists.WithShippingWeightUnitsCreateIsDefault(true),
    shipping_value_lists.WithShippingWeightUnitsCreateLabels(map[string]interface{}{
        "de": "Tonne",
        "en": "Tonne"
    }),
    shipping_value_lists.WithShippingWeightUnitsCreatePosition(1),
    shipping_value_lists.WithShippingWeightUnitsCreateTone("neutral"),
)
```
