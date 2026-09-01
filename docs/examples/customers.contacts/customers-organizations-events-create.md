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

response, error := service.CustomersOrganizationsEventsCreate(
    "",
    "",
    "Called about the annual requirement",
    customers_contacts.WithCustomersOrganizationsEventsCreateActor("vertrieb@example.com"),
    customers_contacts.WithCustomersOrganizationsEventsCreateKind("note"),
    customers_contacts.WithCustomersOrganizationsEventsCreateNote("Asked for a quote on the annual bolt requirement; call back in week 34."),
    customers_contacts.WithCustomersOrganizationsEventsCreateOccurredAt("2026-01-01T12:00:00Z"),
)
```
