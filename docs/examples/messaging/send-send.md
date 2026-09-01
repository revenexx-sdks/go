```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/messaging"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := messaging.New(client)

response, error := service.SendSend(
    "",
    "",
    "",
    messaging.WithSendSendAttachments([]interface{}{}),
    messaging.WithSendSendData(map[string]interface{}{}),
    messaging.WithSendSendDraft(true),
    messaging.WithSendSendLocale(""),
    messaging.WithSendSendMarket(""),
    messaging.WithSendSendSendAt("2026-01-01T12:00:00Z"),
)
```
