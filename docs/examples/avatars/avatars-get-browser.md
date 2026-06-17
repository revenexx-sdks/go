```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/avatars"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com")
    client.WithApiKeyAuth("<API_KEY>")
)

service := avatars.New(client)

response, error := service.AvatarsGetBrowser(
    "",
    avatars.WithAvatarsGetBrowserWidth(0),
    avatars.WithAvatarsGetBrowserHeight(0),
    avatars.WithAvatarsGetBrowserQuality(0),
)
```
