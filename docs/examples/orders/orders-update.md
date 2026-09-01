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

response, error := service.OrdersUpdate(
    "",
    orders.WithOrdersUpdateBillingAddress(map[string]interface{}{
        "city": "Berlin",
        "company": "Beispiel Industrietechnik GmbH",
        "country": "DE",
        "name": "Anna Berger",
        "street": "Musterstra\u00dfe 12",
        "zip": "10115"
    }),
    orders.WithOrdersUpdateBuyer(map[string]interface{}{
        "company": "Beispiel Industrietechnik GmbH",
        "customer_number": "K-10042",
        "email": "anna.berger@example.com",
        "name": "Anna Berger"
    }),
    orders.WithOrdersUpdateCustomerOrderNumber("PO-2026-0042"),
    orders.WithOrdersUpdateMetadata(map[string]interface{}{
        "erp_batch": "2026-W32"
    }),
    orders.WithOrdersUpdateShippingAddress(map[string]interface{}{
        "city": "Berlin",
        "company": "Beispiel Industrietechnik GmbH",
        "country": "DE",
        "name": "Anna Berger",
        "street": "Musterstra\u00dfe 12",
        "zip": "10115"
    }),
    orders.WithOrdersUpdateUserData(map[string]interface{}{
        "campaign": "spring-catalogue",
        "source": "webshop"
    }),
)
```
