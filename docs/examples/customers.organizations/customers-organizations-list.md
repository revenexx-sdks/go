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

response, error := service.CustomersOrganizationsList(
    customers_organizations.WithCustomersOrganizationsListId(""),
    customers_organizations.WithCustomersOrganizationsListName("Beispiel Industrietechnik GmbH"),
    customers_organizations.WithCustomersOrganizationsListVatId("DE123456789"),
    customers_organizations.WithCustomersOrganizationsListBranche("Maschinenbau"),
    customers_organizations.WithCustomersOrganizationsListCustomerNumber("K-10042"),
    customers_organizations.WithCustomersOrganizationsListStatus("active"),
    customers_organizations.WithCustomersOrganizationsListLifecycleStage("customer"),
    customers_organizations.WithCustomersOrganizationsListPaymentTerms("net_30"),
    customers_organizations.WithCustomersOrganizationsListCreditLimit(9.99),
    customers_organizations.WithCustomersOrganizationsListPriceList("standard"),
    customers_organizations.WithCustomersOrganizationsListDeliveryBlock(true),
    customers_organizations.WithCustomersOrganizationsListExternalTeamId(""),
    customers_organizations.WithCustomersOrganizationsListCreatedAt("2026-01-01T12:00:00Z"),
    customers_organizations.WithCustomersOrganizationsListUpdatedAt("2026-01-01T12:00:00Z"),
    customers_organizations.WithCustomersOrganizationsListLimit(1),
    customers_organizations.WithCustomersOrganizationsListOffset(1),
    customers_organizations.WithCustomersOrganizationsListOrder("created_at.desc"),
)
```
