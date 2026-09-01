```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/orders"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := orders.New(client)

response, error := service.OrdersPlace(
    []interface{}{},
    orders.WithOrdersPlaceBillingAddress(map[string]interface{}{
        "city": "Berlin",
        "company": "Beispiel Industrietechnik GmbH",
        "country": "DE",
        "name": "Anna Berger",
        "street": "Musterstra\u00dfe 12",
        "zip": "10115"
    }),
    orders.WithOrdersPlaceBuyer(map[string]interface{}{
        "company": "Beispiel Industrietechnik GmbH",
        "customer_number": "K-10042",
        "email": "anna.berger@example.com",
        "name": "Anna Berger"
    }),
    orders.WithOrdersPlaceCartId(""),
    orders.WithOrdersPlaceChannelId(""),
    orders.WithOrdersPlaceContactId(""),
    orders.WithOrdersPlaceCurrency("EUR"),
    orders.WithOrdersPlaceCustomerOrderNumber("PO-2026-0042"),
    orders.WithOrdersPlaceGrandTotal(243.9),
    orders.WithOrdersPlaceMetadata(map[string]interface{}{
        "erp_batch": "2026-W32"
    }),
    orders.WithOrdersPlaceOrganizationId(""),
    orders.WithOrdersPlacePayment(map[string]interface{}{
        "method": "invoice",
        "status": "open"
    }),
    orders.WithOrdersPlaceShipping(map[string]interface{}{
        "method": "standard",
        "price": 5.9,
        "tax_rate": 19
    }),
    orders.WithOrdersPlaceShippingAddress(map[string]interface{}{
        "city": "Berlin",
        "company": "Beispiel Industrietechnik GmbH",
        "country": "DE",
        "name": "Anna Berger",
        "street": "Musterstra\u00dfe 12",
        "zip": "10115"
    }),
    orders.WithOrdersPlaceShippingTotal(5.9),
    orders.WithOrdersPlaceUserData(map[string]interface{}{
        "campaign": "spring-catalogue",
        "source": "webshop"
    }),
)
```
