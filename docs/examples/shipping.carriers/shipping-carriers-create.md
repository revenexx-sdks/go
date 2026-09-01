```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/shipping_carriers"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := shipping_carriers.New(client)

response, error := service.ShippingCarriersCreate(
    "acme-parcel",
    "Acme Parcel",
    shipping_carriers.WithShippingCarriersCreateCountries(interface{}{"DE","AT","CH"}),
    shipping_carriers.WithShippingCarriersCreateCutoffTime("16:00"),
    shipping_carriers.WithShippingCarriersCreateEtaDaysMax(1),
    shipping_carriers.WithShippingCarriersCreateEtaDaysMin(1),
    shipping_carriers.WithShippingCarriersCreateHandlingDays(1),
    shipping_carriers.WithShippingCarriersCreateLabels(map[string]interface{}{
        "de": "Acme Paketdienst",
        "en": "Acme Parcel"
    }),
    shipping_carriers.WithShippingCarriersCreateMetadata(map[string]interface{}{
        "contract": "ACME-2026",
        "customer_number": "4711"
    }),
    shipping_carriers.WithShippingCarriersCreatePosition(1),
    shipping_carriers.WithShippingCarriersCreateServiceLevel("express"),
    shipping_carriers.WithShippingCarriersCreateStatus("active"),
    shipping_carriers.WithShippingCarriersCreateTrackingUrlTemplate("https://track.example.com/parcels/{tracking_code}"),
)
```
