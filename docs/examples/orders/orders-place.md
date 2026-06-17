```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/orders"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com")
    client.WithApiKeyAuth("<API_KEY>")
)

service := orders.New(client)

response, error := service.OrdersPlace(
    []interface{}{},
    orders.WithOrdersPlaceBillingAddress(map[string]interface{}{}),
    orders.WithOrdersPlaceBuyer(map[string]interface{}{}),
    orders.WithOrdersPlaceCartId(""),
    orders.WithOrdersPlaceChannelId(""),
    orders.WithOrdersPlaceContactId(""),
    orders.WithOrdersPlaceCurrency(""),
    orders.WithOrdersPlaceCustomerOrderNumber(""),
    orders.WithOrdersPlaceGrandTotal(0),
    orders.WithOrdersPlaceMarketId(""),
    orders.WithOrdersPlaceMetadata(map[string]interface{}{}),
    orders.WithOrdersPlaceOrganizationId(""),
    orders.WithOrdersPlacePayment(map[string]interface{}{}),
    orders.WithOrdersPlaceShipping(map[string]interface{}{}),
    orders.WithOrdersPlaceShippingAddress(map[string]interface{}{}),
    orders.WithOrdersPlaceShippingTotal(0),
    orders.WithOrdersPlaceUserData(map[string]interface{}{}),
)
```
