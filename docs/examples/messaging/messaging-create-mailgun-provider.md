```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/messaging"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com")
    client.WithApiKeyAuth("<API_KEY>")
)

service := messaging.New(client)

response, error := service.MessagingCreateMailgunProvider(
    "",
    "",
    messaging.WithMessagingCreateMailgunProviderApiKey(""),
    messaging.WithMessagingCreateMailgunProviderDomain(""),
    messaging.WithMessagingCreateMailgunProviderEnabled(false),
    messaging.WithMessagingCreateMailgunProviderFromEmail(""),
    messaging.WithMessagingCreateMailgunProviderFromName(""),
    messaging.WithMessagingCreateMailgunProviderIsEuRegion(false),
    messaging.WithMessagingCreateMailgunProviderReplyToEmail(""),
    messaging.WithMessagingCreateMailgunProviderReplyToName(""),
)
```
