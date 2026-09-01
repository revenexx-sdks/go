```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/customers_contacts"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := customers_contacts.New(client)

response, error := service.CustomersContactsUpdate(
    "",
    customers_contacts.WithCustomersContactsUpdateEmail("einkauf@example.com"),
    customers_contacts.WithCustomersContactsUpdateFirstName("Anna"),
    customers_contacts.WithCustomersContactsUpdateIsPrimary(true),
    customers_contacts.WithCustomersContactsUpdateJobTitle("Einkaufsleitung"),
    customers_contacts.WithCustomersContactsUpdateLastName("Berger"),
    customers_contacts.WithCustomersContactsUpdateLocale("de-DE"),
    customers_contacts.WithCustomersContactsUpdateOrderApprovalLimit(25000),
    customers_contacts.WithCustomersContactsUpdateOrganizationId(""),
    customers_contacts.WithCustomersContactsUpdatePhone("+49 30 5550123"),
    customers_contacts.WithCustomersContactsUpdateRegistrationStatus("pending"),
    customers_contacts.WithCustomersContactsUpdateRole("buyer"),
    customers_contacts.WithCustomersContactsUpdateStatus("invited"),
)
```
