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

response, error := service.SearchSearchDocuments(
    "",
    search.WithSearchSearchDocumentsFacetBy(""),
    search.WithSearchSearchDocumentsFilterBy(""),
    search.WithSearchSearchDocumentsPage(0),
    search.WithSearchSearchDocumentsPerPage(0),
    search.WithSearchSearchDocumentsQ(""),
    search.WithSearchSearchDocumentsQueryBy(""),
    search.WithSearchSearchDocumentsSortBy(""),
)
```
