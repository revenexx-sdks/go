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

response, error := service.CustomersContactsEventsCreate(
    "",
    "Called about the annual requirement",
    customers_contacts.WithCustomersContactsEventsCreateActor("vertrieb@example.com"),
    customers_contacts.WithCustomersContactsEventsCreateKind("note"),
    customers_contacts.WithCustomersContactsEventsCreateNote("Asked for a quote on the annual bolt requirement; call back in week 34."),
    customers_contacts.WithCustomersContactsEventsCreateOccurredAt("2026-01-01T12:00:00Z"),
)
```
