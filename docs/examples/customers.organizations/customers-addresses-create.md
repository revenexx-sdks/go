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

response, error := service.CustomersAddressesCreate(
    "Berlin",
    "DE",
    "Musterstraße 12",
    "10115",
    customers_organizations.WithCustomersAddressesCreateCompany("Beispiel Industrietechnik GmbH"),
    customers_organizations.WithCustomersAddressesCreateContactId(""),
    customers_organizations.WithCustomersAddressesCreateIsDefault(true),
    customers_organizations.WithCustomersAddressesCreateName("Anna Berger"),
    customers_organizations.WithCustomersAddressesCreateOrganizationId(""),
    customers_organizations.WithCustomersAddressesCreatePhone("+49 30 5550123"),
    customers_organizations.WithCustomersAddressesCreateRegion("Berlin"),
    customers_organizations.WithCustomersAddressesCreateStreet2("Gebäude C, 2. OG"),
    customers_organizations.WithCustomersAddressesCreateType("shipping"),
)
```
