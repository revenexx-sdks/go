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

response, error := service.MessagingUpdateMailgunProvider(
    "",
    messaging.WithMessagingUpdateMailgunProviderApiKey(""),
    messaging.WithMessagingUpdateMailgunProviderDomain(""),
    messaging.WithMessagingUpdateMailgunProviderEnabled(false),
    messaging.WithMessagingUpdateMailgunProviderFromEmail(""),
    messaging.WithMessagingUpdateMailgunProviderFromName(""),
    messaging.WithMessagingUpdateMailgunProviderIsEuRegion(false),
    messaging.WithMessagingUpdateMailgunProviderName(""),
    messaging.WithMessagingUpdateMailgunProviderReplyToEmail(""),
    messaging.WithMessagingUpdateMailgunProviderReplyToName(""),
)
```
