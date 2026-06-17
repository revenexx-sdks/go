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

response, error := service.MessagingUpdatePush(
    "",
    messaging.WithMessagingUpdatePushAction(""),
    messaging.WithMessagingUpdatePushBadge(0),
    messaging.WithMessagingUpdatePushBody(""),
    messaging.WithMessagingUpdatePushColor(""),
    messaging.WithMessagingUpdatePushContentAvailable(false),
    messaging.WithMessagingUpdatePushCritical(false),
    messaging.WithMessagingUpdatePushData(map[string]interface{}{}),
    messaging.WithMessagingUpdatePushDraft(false),
    messaging.WithMessagingUpdatePushIcon(""),
    messaging.WithMessagingUpdatePushImage(""),
    messaging.WithMessagingUpdatePushPriority(""),
    messaging.WithMessagingUpdatePushScheduledAt(""),
    messaging.WithMessagingUpdatePushSound(""),
    messaging.WithMessagingUpdatePushTag(""),
    messaging.WithMessagingUpdatePushTargets([]interface{}{}),
    messaging.WithMessagingUpdatePushTitle(""),
    messaging.WithMessagingUpdatePushTopics([]interface{}{}),
    messaging.WithMessagingUpdatePushUsers([]interface{}{}),
)
```
