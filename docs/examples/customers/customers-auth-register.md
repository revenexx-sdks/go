```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/customers"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := customers.New(client)

response, error := service.CustomersAuthRegister(
    "einkauf@example.com",
    "",
    customers.WithCustomersAuthRegisterFirstName("Anna"),
    customers.WithCustomersAuthRegisterLastName("Berger"),
    customers.WithCustomersAuthRegisterLocale("de-DE"),
    customers.WithCustomersAuthRegisterOrganizationId(""),
    customers.WithCustomersAuthRegisterOrganizationName("Beispiel Industrietechnik GmbH"),
    customers.WithCustomersAuthRegisterUrl("https://shop.example.com/account"),
    customers.WithCustomersAuthRegisterVatId("DE123456789"),
    customers.WithCustomersAuthRegisterVerificationUrl("https://shop.example.com/bestaetigen"),
)
```
