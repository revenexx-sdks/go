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

response, error := service.ShippingCarriersList(
    shipping_carriers.WithShippingCarriersListLimit(1),
    shipping_carriers.WithShippingCarriersListOffset(1),
    shipping_carriers.WithShippingCarriersListOrder("position.asc"),
    shipping_carriers.WithShippingCarriersListCode("acme-parcel"),
    shipping_carriers.WithShippingCarriersListStatus("active"),
    shipping_carriers.WithShippingCarriersListServiceLevel("express"),
)
```
