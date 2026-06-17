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

response, error := service.CustomersContactsCreate(
    "",
    customers.WithCustomersContactsCreateFirstName(""),
    customers.WithCustomersContactsCreateIsPrimary(false),
    customers.WithCustomersContactsCreateLastName(""),
    customers.WithCustomersContactsCreateLocale(""),
    customers.WithCustomersContactsCreateOrganizationId(""),
    customers.WithCustomersContactsCreatePhone(""),
    customers.WithCustomersContactsCreateRole(""),
    customers.WithCustomersContactsCreateStatus(""),
)
```
