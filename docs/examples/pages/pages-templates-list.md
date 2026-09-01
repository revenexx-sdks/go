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

response, error := service.PagesTemplatesList(
    pages.WithPagesTemplatesListLimit(1),
    pages.WithPagesTemplatesListOffset(1),
    pages.WithPagesTemplatesListOrder("created_at.desc"),
    pages.WithPagesTemplatesListId(""),
    pages.WithPagesTemplatesListLabel("Hero with two teasers"),
    pages.WithPagesTemplatesListDescription("Full-width hero followed by a two-column teaser row."),
    pages.WithPagesTemplatesListPageBundle("standard"),
    pages.WithPagesTemplatesListFieldName("content"),
    pages.WithPagesTemplatesListIsDefault(true),
    pages.WithPagesTemplatesListCreatedBy(""),
    pages.WithPagesTemplatesListCreatedAt(""),
    pages.WithPagesTemplatesListUpdatedAt(""),
)
```
