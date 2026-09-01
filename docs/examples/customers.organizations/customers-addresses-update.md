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

response, error := service.CustomersAddressesUpdate(
    "",
    customers_organizations.WithCustomersAddressesUpdateCity("Berlin"),
    customers_organizations.WithCustomersAddressesUpdateCompany("Beispiel Industrietechnik GmbH"),
    customers_organizations.WithCustomersAddressesUpdateContactId(""),
    customers_organizations.WithCustomersAddressesUpdateCountry("DE"),
    customers_organizations.WithCustomersAddressesUpdateIsDefault(true),
    customers_organizations.WithCustomersAddressesUpdateName("Anna Berger"),
    customers_organizations.WithCustomersAddressesUpdateOrganizationId(""),
    customers_organizations.WithCustomersAddressesUpdatePhone("+49 30 5550123"),
    customers_organizations.WithCustomersAddressesUpdateRegion("Berlin"),
    customers_organizations.WithCustomersAddressesUpdateStreet("Musterstraße 12"),
    customers_organizations.WithCustomersAddressesUpdateStreet2("Gebäude C, 2. OG"),
    customers_organizations.WithCustomersAddressesUpdateType("shipping"),
    customers_organizations.WithCustomersAddressesUpdateZip("10115"),
)
```
