```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/pages"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := pages.New(client)

response, error := service.PagesPagesRevisions(
    "",
    pages.WithPagesPagesRevisionsLimit(1),
    pages.WithPagesPagesRevisionsOffset(1),
    pages.WithPagesPagesRevisionsOrder("created_at.desc"),
    pages.WithPagesPagesRevisionsLabel("Autumn campaign"),
    pages.WithPagesPagesRevisionsCreatedBy(""),
    pages.WithPagesPagesRevisionsCreatedByName(""),
    pages.WithPagesPagesRevisionsCreatedAt(""),
)
```
