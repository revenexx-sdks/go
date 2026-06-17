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

response, error := service.MessagingCreateMsg91Provider(
    "",
    "",
    messaging.WithMessagingCreateMsg91ProviderAuthKey(""),
    messaging.WithMessagingCreateMsg91ProviderEnabled(false),
    messaging.WithMessagingCreateMsg91ProviderSenderId(""),
    messaging.WithMessagingCreateMsg91ProviderTemplateId(""),
)
```
