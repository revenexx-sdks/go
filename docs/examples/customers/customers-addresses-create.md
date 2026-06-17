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

response, error := service.CustomersAddressesCreate(
    "",
    "",
    "",
    "",
    customers.WithCustomersAddressesCreateCompany(""),
    customers.WithCustomersAddressesCreateContactId(""),
    customers.WithCustomersAddressesCreateIsDefault(false),
    customers.WithCustomersAddressesCreateName(""),
    customers.WithCustomersAddressesCreateOrganizationId(""),
    customers.WithCustomersAddressesCreatePhone(""),
    customers.WithCustomersAddressesCreateRegion(""),
    customers.WithCustomersAddressesCreateStreet2(""),
    customers.WithCustomersAddressesCreateType(""),
)
```
