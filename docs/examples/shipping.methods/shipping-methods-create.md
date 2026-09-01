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

response, error := service.ShippingMethodsCreate(
    "express",
    "Express delivery",
    shipping_methods.WithShippingMethodsCreateCarrier("acme-parcel"),
    shipping_methods.WithShippingMethodsCreateCarrierId("8a4d1c7e-2b93-4f61-b0d2-6c5a9e3f1a44"),
    shipping_methods.WithShippingMethodsCreateCountries(interface{}{"DE","AT","CH"}),
    shipping_methods.WithShippingMethodsCreateCurrency("EUR"),
    shipping_methods.WithShippingMethodsCreateDescription("Delivered by the next working day when ordered before the cut-off."),
    shipping_methods.WithShippingMethodsCreateEnabled(true),
    shipping_methods.WithShippingMethodsCreateEtaDaysMax(1),
    shipping_methods.WithShippingMethodsCreateEtaDaysMin(1),
    shipping_methods.WithShippingMethodsCreateFreeAbove(100),
    shipping_methods.WithShippingMethodsCreateLabels(map[string]interface{}{
        "de": "Expressversand",
        "en": "Express delivery"
    }),
    shipping_methods.WithShippingMethodsCreateMatrixAttribute("volume_litres"),
    shipping_methods.WithShippingMethodsCreateMatrixBasis("weight"),
    shipping_methods.WithShippingMethodsCreateMetadata(map[string]interface{}{
        "erp_key": "SHIP-EXPRESS",
        "printer": "label-2"
    }),
    shipping_methods.WithShippingMethodsCreatePosition(1),
    shipping_methods.WithShippingMethodsCreatePrice(9.9),
    shipping_methods.WithShippingMethodsCreatePricingType("fixed"),
    shipping_methods.WithShippingMethodsCreateQuoteAbove(31.5),
    shipping_methods.WithShippingMethodsCreateTaxClass("reduced"),
)
```
