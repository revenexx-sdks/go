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

response, error := service.ShippingRates(
    shipping_methods.WithShippingRatesAt("2026-01-01T12:00:00Z"),
    shipping_methods.WithShippingRatesAttributes(map[string]interface{}{
        "volume_litres": 48
    }),
    shipping_methods.WithShippingRatesCountry("DE"),
    shipping_methods.WithShippingRatesCurrency("EUR"),
    shipping_methods.WithShippingRatesMarketId("3f2b6d10-7c41-4c0a-9a35-2f5b8e0d9c11"),
    shipping_methods.WithShippingRatesOrderValue(129.9),
    shipping_methods.WithShippingRatesOrderValueGross(129.9),
    shipping_methods.WithShippingRatesOrderValueNet(109.16),
    shipping_methods.WithShippingRatesQuantity(3),
    shipping_methods.WithShippingRatesWeight(12.5),
    shipping_methods.WithShippingRatesWeightUnit("kg"),
)
```
