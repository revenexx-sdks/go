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

response, error := service.MessagingUpdateSendgridProvider(
    "",
    messaging.WithMessagingUpdateSendgridProviderApiKey(""),
    messaging.WithMessagingUpdateSendgridProviderEnabled(false),
    messaging.WithMessagingUpdateSendgridProviderFromEmail(""),
    messaging.WithMessagingUpdateSendgridProviderFromName(""),
    messaging.WithMessagingUpdateSendgridProviderName(""),
    messaging.WithMessagingUpdateSendgridProviderReplyToEmail(""),
    messaging.WithMessagingUpdateSendgridProviderReplyToName(""),
)
```
