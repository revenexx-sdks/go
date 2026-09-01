```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/forms"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := forms.New(client)

response, error := service.FormsList(
    forms.WithFormsListId(""),
    forms.WithFormsListName("Contact request"),
    forms.WithFormsListSlug("contact"),
    forms.WithFormsListStatus("draft"),
    forms.WithFormsListCreatedAt("2026-01-31T09:15:00Z"),
    forms.WithFormsListUpdatedAt("2026-01-31T09:15:00Z"),
    forms.WithFormsListLimit(50),
    forms.WithFormsListOffset(0),
    forms.WithFormsListOrder("created_at.desc"),
)
```
