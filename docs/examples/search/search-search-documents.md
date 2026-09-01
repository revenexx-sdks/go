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

response, error := service.SearchSearchDocuments(
    "products",
    search.WithSearchSearchDocumentsExcludeFields(""),
    search.WithSearchSearchDocumentsFacetBy(""),
    search.WithSearchSearchDocumentsFilterBy(""),
    search.WithSearchSearchDocumentsGroupBy(""),
    search.WithSearchSearchDocumentsHighlightFullFields(""),
    search.WithSearchSearchDocumentsIncludeFields(""),
    search.WithSearchSearchDocumentsMaxFacetValues(1),
    search.WithSearchSearchDocumentsNumTypos(1),
    search.WithSearchSearchDocumentsPage(1),
    search.WithSearchSearchDocumentsPerPage(1),
    search.WithSearchSearchDocumentsPrefix(""),
    search.WithSearchSearchDocumentsQ(""),
    search.WithSearchSearchDocumentsQueryBy(""),
    search.WithSearchSearchDocumentsSortBy(""),
)
```
