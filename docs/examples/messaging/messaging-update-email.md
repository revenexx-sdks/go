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

response, error := service.MessagingUpdateEmail(
    "",
    messaging.WithMessagingUpdateEmailAttachments([]interface{}{}),
    messaging.WithMessagingUpdateEmailBcc([]interface{}{}),
    messaging.WithMessagingUpdateEmailCc([]interface{}{}),
    messaging.WithMessagingUpdateEmailContent(""),
    messaging.WithMessagingUpdateEmailDraft(false),
    messaging.WithMessagingUpdateEmailHtml(false),
    messaging.WithMessagingUpdateEmailScheduledAt(""),
    messaging.WithMessagingUpdateEmailSubject(""),
    messaging.WithMessagingUpdateEmailTargets([]interface{}{}),
    messaging.WithMessagingUpdateEmailTopics([]interface{}{}),
    messaging.WithMessagingUpdateEmailUsers([]interface{}{}),
)
```
