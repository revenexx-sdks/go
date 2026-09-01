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

response, error := service.CustomersContactsCreate(
    "einkauf@example.com",
    customers_contacts.WithCustomersContactsCreateFirstName("Anna"),
    customers_contacts.WithCustomersContactsCreateIsPrimary(true),
    customers_contacts.WithCustomersContactsCreateJobTitle("Einkaufsleitung"),
    customers_contacts.WithCustomersContactsCreateLastName("Berger"),
    customers_contacts.WithCustomersContactsCreateLocale("de-DE"),
    customers_contacts.WithCustomersContactsCreateOrderApprovalLimit(25000),
    customers_contacts.WithCustomersContactsCreateOrganizationId(""),
    customers_contacts.WithCustomersContactsCreatePhone("+49 30 5550123"),
    customers_contacts.WithCustomersContactsCreateRegistrationStatus("pending"),
    customers_contacts.WithCustomersContactsCreateRole("buyer"),
    customers_contacts.WithCustomersContactsCreateStatus("invited"),
)
```
