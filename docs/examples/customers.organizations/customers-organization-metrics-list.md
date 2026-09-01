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

response, error := service.CustomersOrganizationMetricsList(
    customers_organizations.WithCustomersOrganizationMetricsListId(""),
    customers_organizations.WithCustomersOrganizationMetricsListOrganizationId(""),
    customers_organizations.WithCustomersOrganizationMetricsListOrderCount(1),
    customers_organizations.WithCustomersOrganizationMetricsListOrderCount30d(1),
    customers_organizations.WithCustomersOrganizationMetricsListOrderCount90d(1),
    customers_organizations.WithCustomersOrganizationMetricsListOrderCount365d(1),
    customers_organizations.WithCustomersOrganizationMetricsListRevenueTotal(9.99),
    customers_organizations.WithCustomersOrganizationMetricsListRevenue30d(9.99),
    customers_organizations.WithCustomersOrganizationMetricsListRevenue90d(9.99),
    customers_organizations.WithCustomersOrganizationMetricsListRevenue365d(9.99),
    customers_organizations.WithCustomersOrganizationMetricsListAvgOrderValue(9.99),
    customers_organizations.WithCustomersOrganizationMetricsListAvgOrderValue365d(9.99),
    customers_organizations.WithCustomersOrganizationMetricsListFirstOrderAt("2026-01-01T12:00:00Z"),
    customers_organizations.WithCustomersOrganizationMetricsListLastOrderAt("2026-01-01T12:00:00Z"),
    customers_organizations.WithCustomersOrganizationMetricsListCurrency("EUR"),
    customers_organizations.WithCustomersOrganizationMetricsListCurrencyMixed(true),
    customers_organizations.WithCustomersOrganizationMetricsListOrdersAsOf("2026-01-01T12:00:00Z"),
    customers_organizations.WithCustomersOrganizationMetricsListComputedAt("2026-01-01T12:00:00Z"),
    customers_organizations.WithCustomersOrganizationMetricsListCreatedAt("2026-01-01T12:00:00Z"),
    customers_organizations.WithCustomersOrganizationMetricsListUpdatedAt("2026-01-01T12:00:00Z"),
    customers_organizations.WithCustomersOrganizationMetricsListLimit(1),
    customers_organizations.WithCustomersOrganizationMetricsListOffset(1),
    customers_organizations.WithCustomersOrganizationMetricsListOrder("created_at.desc"),
)
```
