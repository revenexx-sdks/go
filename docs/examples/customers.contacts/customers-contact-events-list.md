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

response, error := service.CustomersContactEventsList(
    customers_contacts.WithCustomersContactEventsListId(""),
    customers_contacts.WithCustomersContactEventsListContactId(""),
    customers_contacts.WithCustomersContactEventsListOrganizationId(""),
    customers_contacts.WithCustomersContactEventsListKind("call"),
    customers_contacts.WithCustomersContactEventsListName("activity.call"),
    customers_contacts.WithCustomersContactEventsListSubject("Called about the annual requirement"),
    customers_contacts.WithCustomersContactEventsListActor("vertrieb@example.com"),
    customers_contacts.WithCustomersContactEventsListOccurredAt("2026-01-01T12:00:00Z"),
    customers_contacts.WithCustomersContactEventsListCreatedAt("2026-01-01T12:00:00Z"),
    customers_contacts.WithCustomersContactEventsListLimit(1),
    customers_contacts.WithCustomersContactEventsListOffset(1),
    customers_contacts.WithCustomersContactEventsListOrder("created_at.desc"),
)
```
