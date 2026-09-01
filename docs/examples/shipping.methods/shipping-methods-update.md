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

response, error := service.ShippingMethodsUpdate(
    "",
    shipping_methods.WithShippingMethodsUpdateCarrier("acme-parcel"),
    shipping_methods.WithShippingMethodsUpdateCarrierId("8a4d1c7e-2b93-4f61-b0d2-6c5a9e3f1a44"),
    shipping_methods.WithShippingMethodsUpdateCode("express"),
    shipping_methods.WithShippingMethodsUpdateCountries(interface{}{"DE","AT","CH"}),
    shipping_methods.WithShippingMethodsUpdateCurrency("EUR"),
    shipping_methods.WithShippingMethodsUpdateDescription("Delivered by the next working day when ordered before the cut-off."),
    shipping_methods.WithShippingMethodsUpdateEnabled(true),
    shipping_methods.WithShippingMethodsUpdateEtaDaysMax(1),
    shipping_methods.WithShippingMethodsUpdateEtaDaysMin(1),
    shipping_methods.WithShippingMethodsUpdateFreeAbove(100),
    shipping_methods.WithShippingMethodsUpdateLabels(map[string]interface{}{
        "de": "Expressversand",
        "en": "Express delivery"
    }),
    shipping_methods.WithShippingMethodsUpdateMatrixAttribute("volume_litres"),
    shipping_methods.WithShippingMethodsUpdateMatrixBasis("weight"),
    shipping_methods.WithShippingMethodsUpdateMetadata(map[string]interface{}{
        "erp_key": "SHIP-EXPRESS",
        "printer": "label-2"
    }),
    shipping_methods.WithShippingMethodsUpdateName("Express delivery"),
    shipping_methods.WithShippingMethodsUpdatePosition(1),
    shipping_methods.WithShippingMethodsUpdatePrice(9.9),
    shipping_methods.WithShippingMethodsUpdatePricingType("fixed"),
    shipping_methods.WithShippingMethodsUpdateQuoteAbove(31.5),
    shipping_methods.WithShippingMethodsUpdateTaxClass("reduced"),
)
```
