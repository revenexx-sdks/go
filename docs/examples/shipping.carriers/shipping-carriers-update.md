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

response, error := service.ShippingCarriersUpdate(
    "",
    shipping_carriers.WithShippingCarriersUpdateCode("acme-parcel"),
    shipping_carriers.WithShippingCarriersUpdateCountries(interface{}{"DE","AT","CH"}),
    shipping_carriers.WithShippingCarriersUpdateCutoffTime("16:00"),
    shipping_carriers.WithShippingCarriersUpdateEtaDaysMax(1),
    shipping_carriers.WithShippingCarriersUpdateEtaDaysMin(1),
    shipping_carriers.WithShippingCarriersUpdateHandlingDays(1),
    shipping_carriers.WithShippingCarriersUpdateLabels(map[string]interface{}{
        "de": "Acme Paketdienst",
        "en": "Acme Parcel"
    }),
    shipping_carriers.WithShippingCarriersUpdateMetadata(map[string]interface{}{
        "contract": "ACME-2026",
        "customer_number": "4711"
    }),
    shipping_carriers.WithShippingCarriersUpdateName("Acme Parcel"),
    shipping_carriers.WithShippingCarriersUpdatePosition(1),
    shipping_carriers.WithShippingCarriersUpdateServiceLevel("express"),
    shipping_carriers.WithShippingCarriersUpdateStatus("active"),
    shipping_carriers.WithShippingCarriersUpdateTrackingUrlTemplate("https://track.example.com/parcels/{tracking_code}"),
)
```
