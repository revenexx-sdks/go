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

response, error := service.ShippingServiceLevelsCreate(
    "night_courier",
    "Night courier",
    shipping_value_lists.WithShippingServiceLevelsCreateDescription("When to pick this service level."),
    shipping_value_lists.WithShippingServiceLevelsCreateDescriptions(map[string]interface{}{
        "de": "Wann diese Option zu w\u00e4hlen ist.",
        "en": "When to pick this service level."
    }),
    shipping_value_lists.WithShippingServiceLevelsCreateIsDefault(true),
    shipping_value_lists.WithShippingServiceLevelsCreateLabels(map[string]interface{}{
        "de": "Night courier",
        "en": "Night courier"
    }),
    shipping_value_lists.WithShippingServiceLevelsCreatePosition(1),
    shipping_value_lists.WithShippingServiceLevelsCreateTone("neutral"),
)
```
