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

response, error := service.MessagingUpdateMsg91Provider(
    "",
    messaging.WithMessagingUpdateMsg91ProviderAuthKey(""),
    messaging.WithMessagingUpdateMsg91ProviderEnabled(false),
    messaging.WithMessagingUpdateMsg91ProviderName(""),
    messaging.WithMessagingUpdateMsg91ProviderSenderId(""),
    messaging.WithMessagingUpdateMsg91ProviderTemplateId(""),
)
```
