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

response, error := service.BindingUpdate(
    "",
    messaging.WithBindingUpdateChannel(""),
    messaging.WithBindingUpdateEnabled(true),
    messaging.WithBindingUpdateEventTopic(""),
    messaging.WithBindingUpdateFallbackOrder(1),
    messaging.WithBindingUpdateLocale(""),
    messaging.WithBindingUpdateRecipient(""),
    messaging.WithBindingUpdateTemplateKey(""),
)
```
