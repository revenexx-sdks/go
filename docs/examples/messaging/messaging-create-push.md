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

response, error := service.MessagingCreatePush(
    "",
    messaging.WithMessagingCreatePushAction(""),
    messaging.WithMessagingCreatePushBadge(0),
    messaging.WithMessagingCreatePushBody(""),
    messaging.WithMessagingCreatePushColor(""),
    messaging.WithMessagingCreatePushContentAvailable(false),
    messaging.WithMessagingCreatePushCritical(false),
    messaging.WithMessagingCreatePushData(map[string]interface{}{}),
    messaging.WithMessagingCreatePushDraft(false),
    messaging.WithMessagingCreatePushIcon(""),
    messaging.WithMessagingCreatePushImage(""),
    messaging.WithMessagingCreatePushPriority(""),
    messaging.WithMessagingCreatePushScheduledAt(""),
    messaging.WithMessagingCreatePushSound(""),
    messaging.WithMessagingCreatePushTag(""),
    messaging.WithMessagingCreatePushTargets([]interface{}{}),
    messaging.WithMessagingCreatePushTitle(""),
    messaging.WithMessagingCreatePushTopics([]interface{}{}),
    messaging.WithMessagingCreatePushUsers([]interface{}{}),
)
```
