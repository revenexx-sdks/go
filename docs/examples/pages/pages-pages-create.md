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

response, error := service.PagesPagesCreate(
    "About us",
    pages.WithPagesPagesCreateBundle("standard"),
    pages.WithPagesPagesCreateHostOptions(map[string]interface{}{}),
    pages.WithPagesPagesCreateMeta(map[string]interface{}{}),
    pages.WithPagesPagesCreateSlug("about-us"),
    pages.WithPagesPagesCreateSourceLanguage("de"),
)
```
