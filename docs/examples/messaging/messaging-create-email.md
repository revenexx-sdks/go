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

response, error := service.MessagingCreateEmail(
    "",
    "",
    "",
    messaging.WithMessagingCreateEmailAttachments([]interface{}{}),
    messaging.WithMessagingCreateEmailBcc([]interface{}{}),
    messaging.WithMessagingCreateEmailCc([]interface{}{}),
    messaging.WithMessagingCreateEmailDraft(false),
    messaging.WithMessagingCreateEmailHtml(false),
    messaging.WithMessagingCreateEmailScheduledAt(""),
    messaging.WithMessagingCreateEmailTargets([]interface{}{}),
    messaging.WithMessagingCreateEmailTopics([]interface{}{}),
    messaging.WithMessagingCreateEmailUsers([]interface{}{}),
)
```
