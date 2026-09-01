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

response, error := service.PagesPagesUpdate(
    "",
    pages.WithPagesPagesUpdateBundle("standard"),
    pages.WithPagesPagesUpdateMeta(map[string]interface{}{}),
    pages.WithPagesPagesUpdateSlug("about-us"),
    pages.WithPagesPagesUpdateStatus("draft"),
    pages.WithPagesPagesUpdateTitle("About us"),
)
```
