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

response, error := service.PagesPagesList(
    pages.WithPagesPagesListLimit(1),
    pages.WithPagesPagesListOffset(1),
    pages.WithPagesPagesListOrder("created_at.desc"),
    pages.WithPagesPagesListBundle("standard"),
    pages.WithPagesPagesListStatus("draft"),
    pages.WithPagesPagesListQ("contact"),
)
```
