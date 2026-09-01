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

response, error := service.OrdersItemsCancel(
    "",
    []interface{}{},
    orders.WithOrdersItemsCancelCancelledBy("service-desk"),
    orders.WithOrdersItemsCancelReason("Out of stock, customer agreed"),
)
```
