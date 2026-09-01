```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/pages_editor"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := pages_editor.New(client)

response, error := service.PagesEditorTemplatesCreate(
    "",
    "Hero with two teasers",
    []interface{}{},
    pages_editor.WithPagesEditorTemplatesCreateDescription("Full-width hero followed by a two-column teaser row."),
    pages_editor.WithPagesEditorTemplatesCreateFieldName("content"),
    pages_editor.WithPagesEditorTemplatesCreateIsDefault(true),
    pages_editor.WithPagesEditorTemplatesCreatePageBundle("standard"),
)
```
