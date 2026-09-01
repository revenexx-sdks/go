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

response, error := service.PagesLibraryList(
    pages.WithPagesLibraryListLimit(1),
    pages.WithPagesLibraryListOffset(1),
    pages.WithPagesLibraryListOrder("created_at.desc"),
    pages.WithPagesLibraryListBundles("hero,teaser"),
    pages.WithPagesLibraryListText("hero"),
)
```
