```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/pages"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com")
    client.WithApiKeyAuth("<API_KEY>")
)

service := pages.New(client)

response, error := service.PagesPagesCreate(
    "",
    pages.WithPagesPagesCreateBundle(""),
    pages.WithPagesPagesCreateHostOptions(map[string]interface{}{}),
    pages.WithPagesPagesCreateMeta(map[string]interface{}{}),
    pages.WithPagesPagesCreateSlug(""),
    pages.WithPagesPagesCreateSourceLanguage(""),
)
```
