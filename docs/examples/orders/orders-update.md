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

response, error := service.OrdersUpdate(
    "",
    orders.WithOrdersUpdateBillingAddress(map[string]interface{}{}),
    orders.WithOrdersUpdateBuyer(map[string]interface{}{}),
    orders.WithOrdersUpdateCustomerOrderNumber(""),
    orders.WithOrdersUpdateMetadata(map[string]interface{}{}),
    orders.WithOrdersUpdateShippingAddress(map[string]interface{}{}),
    orders.WithOrdersUpdateUserData(map[string]interface{}{}),
)
```
