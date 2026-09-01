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

response, error := service.CustomersOrganizationsCreate(
    "Beispiel Industrietechnik GmbH",
    customers_organizations.WithCustomersOrganizationsCreateBranche("Maschinenbau"),
    customers_organizations.WithCustomersOrganizationsCreateCreditLimit(5000),
    customers_organizations.WithCustomersOrganizationsCreateCustomerNumber("K-10042"),
    customers_organizations.WithCustomersOrganizationsCreateDeliveryBlock(true),
    customers_organizations.WithCustomersOrganizationsCreateLifecycleStage("customer"),
    customers_organizations.WithCustomersOrganizationsCreatePaymentTerms("net_30"),
    customers_organizations.WithCustomersOrganizationsCreatePriceList("standard"),
    customers_organizations.WithCustomersOrganizationsCreateSettings(map[string]interface{}{
        "account_manager": "sales-north",
        "delivery_tour": "tuesday",
        "self_pickup": true
    }),
    customers_organizations.WithCustomersOrganizationsCreateStatus("active"),
    customers_organizations.WithCustomersOrganizationsCreateVatId("DE123456789"),
)
```
