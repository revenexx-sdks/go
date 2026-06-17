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

response, error := service.AvatarsGetCreditCard(
    "",
    avatars.WithAvatarsGetCreditCardWidth(0),
    avatars.WithAvatarsGetCreditCardHeight(0),
    avatars.WithAvatarsGetCreditCardQuality(0),
)
```
