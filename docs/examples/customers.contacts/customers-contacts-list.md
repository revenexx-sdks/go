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

response, error := service.CustomersContactsList(
    customers_contacts.WithCustomersContactsListId(""),
    customers_contacts.WithCustomersContactsListOrganizationId(""),
    customers_contacts.WithCustomersContactsListEmail("einkauf@example.com"),
    customers_contacts.WithCustomersContactsListFirstName("Anna"),
    customers_contacts.WithCustomersContactsListLastName("Berger"),
    customers_contacts.WithCustomersContactsListPhone("+49 30 5550123"),
    customers_contacts.WithCustomersContactsListJobTitle("Einkaufsleitung"),
    customers_contacts.WithCustomersContactsListRole("buyer"),
    customers_contacts.WithCustomersContactsListStatus("invited"),
    customers_contacts.WithCustomersContactsListOrderApprovalLimit(9.99),
    customers_contacts.WithCustomersContactsListRegistrationStatus("pending"),
    customers_contacts.WithCustomersContactsListRegistrationDecidedAt("2026-01-01T12:00:00Z"),
    customers_contacts.WithCustomersContactsListRegistrationDecidedBy("vertrieb@example.com"),
    customers_contacts.WithCustomersContactsListRegistrationReason("Could not be verified as a commercial buyer."),
    customers_contacts.WithCustomersContactsListLocale("de-DE"),
    customers_contacts.WithCustomersContactsListIsPrimary(true),
    customers_contacts.WithCustomersContactsListExternalUserId(""),
    customers_contacts.WithCustomersContactsListCreatedAt("2026-01-01T12:00:00Z"),
    customers_contacts.WithCustomersContactsListUpdatedAt("2026-01-01T12:00:00Z"),
    customers_contacts.WithCustomersContactsListLimit(1),
    customers_contacts.WithCustomersContactsListOffset(1),
    customers_contacts.WithCustomersContactsListOrder("created_at.desc"),
)
```
