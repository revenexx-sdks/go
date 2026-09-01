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

response, error := service.AvatarsGetCreditCard(
    "amex",
    avatars.WithAvatarsGetCreditCardWidth(1),
    avatars.WithAvatarsGetCreditCardHeight(1),
    avatars.WithAvatarsGetCreditCardQuality(1),
)
```
