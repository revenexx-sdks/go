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

response, error := service.ShippingMethodsList(
    shipping_methods.WithShippingMethodsListLimit(1),
    shipping_methods.WithShippingMethodsListOffset(1),
    shipping_methods.WithShippingMethodsListOrder("position.asc"),
    shipping_methods.WithShippingMethodsListCode("express"),
    shipping_methods.WithShippingMethodsListEnabled(true),
    shipping_methods.WithShippingMethodsListPricingType("matrix"),
    shipping_methods.WithShippingMethodsListCarrierId("8a4d1c7e-2b93-4f61-b0d2-6c5a9e3f1a44"),
    shipping_methods.WithShippingMethodsListCarrier("acme-parcel"),
    shipping_methods.WithShippingMethodsListTaxClass("reduced"),
)
```
