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

response, error := service.CustomersOrganizationsUpdate(
    "",
    customers.WithCustomersOrganizationsUpdateName(""),
    customers.WithCustomersOrganizationsUpdateSettings(map[string]interface{}{}),
    customers.WithCustomersOrganizationsUpdateStatus(""),
    customers.WithCustomersOrganizationsUpdateVatId(""),
)
```
