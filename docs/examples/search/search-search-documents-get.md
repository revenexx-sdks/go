```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/search"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com")
    client.WithApiKeyAuth("<API_KEY>")
)

service := search.New(client)

response, error := service.SearchSearchDocumentsGet(
    "",
    search.WithSearchSearchDocumentsGetQ(""),
    search.WithSearchSearchDocumentsGetQueryBy(""),
    search.WithSearchSearchDocumentsGetFilterBy(""),
    search.WithSearchSearchDocumentsGetSortBy(""),
    search.WithSearchSearchDocumentsGetPage(0),
    search.WithSearchSearchDocumentsGetPerPage(0),
)
```
