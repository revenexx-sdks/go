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

response, error := service.ShippingMethodsCreate(
    "",
    "",
    shipping.WithShippingMethodsCreateCarrier(""),
    shipping.WithShippingMethodsCreateCountries([]interface{}{}),
    shipping.WithShippingMethodsCreateCurrency(""),
    shipping.WithShippingMethodsCreateDescription(""),
    shipping.WithShippingMethodsCreateEnabled(false),
    shipping.WithShippingMethodsCreateEtaDaysMax(0),
    shipping.WithShippingMethodsCreateEtaDaysMin(0),
    shipping.WithShippingMethodsCreateFreeAbove(0),
    shipping.WithShippingMethodsCreateLabels(map[string]interface{}{}),
    shipping.WithShippingMethodsCreateMatrixAttribute(""),
    shipping.WithShippingMethodsCreateMatrixBasis(""),
    shipping.WithShippingMethodsCreateMetadata(map[string]interface{}{}),
    shipping.WithShippingMethodsCreatePosition(0),
    shipping.WithShippingMethodsCreatePrice(0),
    shipping.WithShippingMethodsCreatePricingType(""),
)
```
