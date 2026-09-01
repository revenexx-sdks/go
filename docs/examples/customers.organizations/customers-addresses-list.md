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

response, error := service.CustomersAddressesList(
    customers_organizations.WithCustomersAddressesListId(""),
    customers_organizations.WithCustomersAddressesListOrganizationId(""),
    customers_organizations.WithCustomersAddressesListContactId(""),
    customers_organizations.WithCustomersAddressesListType("shipping"),
    customers_organizations.WithCustomersAddressesListCompany("Beispiel Industrietechnik GmbH"),
    customers_organizations.WithCustomersAddressesListName("Anna Berger"),
    customers_organizations.WithCustomersAddressesListStreet("Musterstraße 12"),
    customers_organizations.WithCustomersAddressesListStreet2("Gebäude C, 2. OG"),
    customers_organizations.WithCustomersAddressesListZip("10115"),
    customers_organizations.WithCustomersAddressesListCity("Berlin"),
    customers_organizations.WithCustomersAddressesListRegion("Berlin"),
    customers_organizations.WithCustomersAddressesListCountry("DE"),
    customers_organizations.WithCustomersAddressesListPhone("+49 30 5550123"),
    customers_organizations.WithCustomersAddressesListIsDefault(true),
    customers_organizations.WithCustomersAddressesListCreatedAt("2026-01-01T12:00:00Z"),
    customers_organizations.WithCustomersAddressesListUpdatedAt("2026-01-01T12:00:00Z"),
    customers_organizations.WithCustomersAddressesListLimit(1),
    customers_organizations.WithCustomersAddressesListOffset(1),
    customers_organizations.WithCustomersAddressesListOrder("created_at.desc"),
)
```
