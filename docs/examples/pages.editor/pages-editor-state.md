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

response, error := service.PagesEditorState(
    "",
    pages_editor.WithPagesEditorStateLangcode("de"),
    pages_editor.WithPagesEditorStateIndex(1),
)
```
