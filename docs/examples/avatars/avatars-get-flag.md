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

response, error := service.AvatarsGetFlag(
    "",
    avatars.WithAvatarsGetFlagWidth(0),
    avatars.WithAvatarsGetFlagHeight(0),
    avatars.WithAvatarsGetFlagQuality(0),
)
```
