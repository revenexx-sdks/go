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

response, error := service.PagesTemplatesUpdate(
    "",
    pages.WithPagesTemplatesUpdateDescription("Full-width hero followed by a two-column teaser row."),
    pages.WithPagesTemplatesUpdateFieldName("content"),
    pages.WithPagesTemplatesUpdateIsDefault(true),
    pages.WithPagesTemplatesUpdateLabel("Hero with two teasers"),
    pages.WithPagesTemplatesUpdatePageBundle("standard"),
    pages.WithPagesTemplatesUpdateTree([]interface{}{}),
)
```
