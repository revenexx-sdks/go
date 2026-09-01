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

response, error := service.FormsSubmissionsList(
    forms.WithFormsSubmissionsListId(""),
    forms.WithFormsSubmissionsListFormId(""),
    forms.WithFormsSubmissionsListFormSlug("contact"),
    forms.WithFormsSubmissionsListSource("/contact"),
    forms.WithFormsSubmissionsListStatus("new"),
    forms.WithFormsSubmissionsListCreatedAt("2026-01-31T09:15:00Z"),
    forms.WithFormsSubmissionsListUpdatedAt("2026-01-31T09:15:00Z"),
    forms.WithFormsSubmissionsListLimit(50),
    forms.WithFormsSubmissionsListOffset(0),
    forms.WithFormsSubmissionsListOrder("created_at.desc"),
)
```
