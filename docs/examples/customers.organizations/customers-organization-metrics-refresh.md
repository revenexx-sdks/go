```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/customers_organizations"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := customers_organizations.New(client)

response, error := service.CustomersOrganizationMetricsRefresh(
    customers_organizations.WithCustomersOrganizationMetricsRefreshAsOf("2026-01-01T12:00:00Z"),
    customers_organizations.WithCustomersOrganizationMetricsRefreshCursor(""),
    customers_organizations.WithCustomersOrganizationMetricsRefreshOrganizationIds([]interface{}{}),
)
```
