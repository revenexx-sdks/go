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

response, error := service.OrdersReportsCustomerRollup(
    orders.WithOrdersReportsCustomerRollupAsOf("2026-01-01T12:00:00Z"),
    orders.WithOrdersReportsCustomerRollupCursor(""),
    orders.WithOrdersReportsCustomerRollupOrganizationIds([]interface{}{}),
    orders.WithOrdersReportsCustomerRollupStatuses([]interface{}{}),
)
```
