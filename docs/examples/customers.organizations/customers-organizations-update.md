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

response, error := service.CustomersOrganizationsUpdate(
    "",
    customers_organizations.WithCustomersOrganizationsUpdateBranche("Maschinenbau"),
    customers_organizations.WithCustomersOrganizationsUpdateCreditLimit(5000),
    customers_organizations.WithCustomersOrganizationsUpdateCustomerNumber("K-10042"),
    customers_organizations.WithCustomersOrganizationsUpdateDeliveryBlock(true),
    customers_organizations.WithCustomersOrganizationsUpdateLifecycleStage("customer"),
    customers_organizations.WithCustomersOrganizationsUpdateName("Beispiel Industrietechnik GmbH"),
    customers_organizations.WithCustomersOrganizationsUpdatePaymentTerms("net_30"),
    customers_organizations.WithCustomersOrganizationsUpdatePriceList("standard"),
    customers_organizations.WithCustomersOrganizationsUpdateSettings(map[string]interface{}{
        "account_manager": "sales-north",
        "delivery_tour": "tuesday",
        "self_pickup": true
    }),
    customers_organizations.WithCustomersOrganizationsUpdateStatus("active"),
    customers_organizations.WithCustomersOrganizationsUpdateVatId("DE123456789"),
)
```
