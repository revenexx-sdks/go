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

response, error := service.AvatarsGetFlag(
    "af",
    avatars.WithAvatarsGetFlagWidth(1),
    avatars.WithAvatarsGetFlagHeight(1),
    avatars.WithAvatarsGetFlagQuality(1),
)
```
