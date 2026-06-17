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

response, error := service.ShippingRates(
    shipping.WithShippingRatesAttributes(map[string]interface{}{}),
    shipping.WithShippingRatesCountry(""),
    shipping.WithShippingRatesCurrency(""),
    shipping.WithShippingRatesMarketId(""),
    shipping.WithShippingRatesOrderValue(0),
    shipping.WithShippingRatesQuantity(0),
    shipping.WithShippingRatesWeight(0),
)
```
