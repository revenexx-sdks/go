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

response, error := service.FormsSubmissionsPrune(
    forms.WithFormsSubmissionsPruneDryRun(true),
    forms.WithFormsSubmissionsPruneFormSlug("contact"),
    forms.WithFormsSubmissionsPruneOlderThanDays(1),
    forms.WithFormsSubmissionsPruneStatus("new"),
)
```
