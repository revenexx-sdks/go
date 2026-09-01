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

response, error := service.PagesMenusList(
    pages.WithPagesMenusListLimit(1),
    pages.WithPagesMenusListOffset(1),
    pages.WithPagesMenusListOrder("created_at.desc"),
)
```
