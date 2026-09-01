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

response, error := service.AvatarsGetInitials(
    avatars.WithAvatarsGetInitialsName("Ada Lovelace"),
    avatars.WithAvatarsGetInitialsWidth(1),
    avatars.WithAvatarsGetInitialsHeight(1),
    avatars.WithAvatarsGetInitialsBackground("1a73e8"),
)
```
