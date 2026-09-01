```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/search"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := search.New(client)

response, error := service.SearchSearchDocumentsGet(
    "products",
    search.WithSearchSearchDocumentsGetQ(""),
    search.WithSearchSearchDocumentsGetQueryBy(""),
    search.WithSearchSearchDocumentsGetFilterBy(""),
    search.WithSearchSearchDocumentsGetSortBy(""),
    search.WithSearchSearchDocumentsGetFacetBy(""),
    search.WithSearchSearchDocumentsGetMaxFacetValues(1),
    search.WithSearchSearchDocumentsGetGroupBy(""),
    search.WithSearchSearchDocumentsGetIncludeFields(""),
    search.WithSearchSearchDocumentsGetExcludeFields(""),
    search.WithSearchSearchDocumentsGetHighlightFullFields(""),
    search.WithSearchSearchDocumentsGetNumTypos(1),
    search.WithSearchSearchDocumentsGetPrefix(""),
    search.WithSearchSearchDocumentsGetPage(1),
    search.WithSearchSearchDocumentsGetPerPage(1),
)
```
