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

response, error := service.PagesPagesUpdate(
    "",
    pages.WithPagesPagesUpdateBundle(""),
    pages.WithPagesPagesUpdateMeta(map[string]interface{}{}),
    pages.WithPagesPagesUpdateSlug(""),
    pages.WithPagesPagesUpdateStatus(""),
    pages.WithPagesPagesUpdateTitle(""),
)
```
