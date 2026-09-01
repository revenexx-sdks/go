```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/avatars"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := avatars.New(client)

response, error := service.AvatarsGetImage(
    "https://www.revenexx.com/img/hero-revenexx-poster.webp",
    avatars.WithAvatarsGetImageWidth(1),
    avatars.WithAvatarsGetImageHeight(1),
)
```
