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

response, error := service.ShippingMethodsUpdate(
    "",
    shipping.WithShippingMethodsUpdateCarrier(""),
    shipping.WithShippingMethodsUpdateCode(""),
    shipping.WithShippingMethodsUpdateCountries([]interface{}{}),
    shipping.WithShippingMethodsUpdateCurrency(""),
    shipping.WithShippingMethodsUpdateDescription(""),
    shipping.WithShippingMethodsUpdateEnabled(false),
    shipping.WithShippingMethodsUpdateEtaDaysMax(0),
    shipping.WithShippingMethodsUpdateEtaDaysMin(0),
    shipping.WithShippingMethodsUpdateFreeAbove(0),
    shipping.WithShippingMethodsUpdateLabels(map[string]interface{}{}),
    shipping.WithShippingMethodsUpdateMatrixAttribute(""),
    shipping.WithShippingMethodsUpdateMatrixBasis(""),
    shipping.WithShippingMethodsUpdateMetadata(map[string]interface{}{}),
    shipping.WithShippingMethodsUpdateName(""),
    shipping.WithShippingMethodsUpdatePosition(0),
    shipping.WithShippingMethodsUpdatePrice(0),
    shipping.WithShippingMethodsUpdatePricingType(""),
)
```
