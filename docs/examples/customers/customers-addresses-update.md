```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/customers"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com")
    client.WithApiKeyAuth("<API_KEY>")
)

service := customers.New(client)

response, error := service.CustomersAddressesUpdate(
    "",
    customers.WithCustomersAddressesUpdateCity(""),
    customers.WithCustomersAddressesUpdateCompany(""),
    customers.WithCustomersAddressesUpdateContactId(""),
    customers.WithCustomersAddressesUpdateCountry(""),
    customers.WithCustomersAddressesUpdateIsDefault(false),
    customers.WithCustomersAddressesUpdateName(""),
    customers.WithCustomersAddressesUpdateOrganizationId(""),
    customers.WithCustomersAddressesUpdatePhone(""),
    customers.WithCustomersAddressesUpdateRegion(""),
    customers.WithCustomersAddressesUpdateStreet(""),
    customers.WithCustomersAddressesUpdateStreet2(""),
    customers.WithCustomersAddressesUpdateType(""),
    customers.WithCustomersAddressesUpdateZip(""),
)
```
