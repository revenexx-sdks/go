```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/pages_collaboration"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := pages_collaboration.New(client)

response, error := service.PagesEditorNotificationsUnreadCount()
```
