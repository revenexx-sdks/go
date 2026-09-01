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

response, error := service.TemplateUpdate(
    "",
    messaging.WithTemplateUpdateBodyHtml(""),
    messaging.WithTemplateUpdateBodyText(""),
    messaging.WithTemplateUpdateContentSid(""),
    messaging.WithTemplateUpdateDesign([]interface{}{}),
    messaging.WithTemplateUpdateEnabled(true),
    messaging.WithTemplateUpdateLayoutId(""),
    messaging.WithTemplateUpdateMarkets([]interface{}{}),
    messaging.WithTemplateUpdateMessageClass("transactional"),
    messaging.WithTemplateUpdateSubject(""),
    messaging.WithTemplateUpdateTestMode(true),
    messaging.WithTemplateUpdateTitle(""),
    messaging.WithTemplateUpdateValidFrom("2026-01-01T12:00:00Z"),
    messaging.WithTemplateUpdateValidUntil("2026-01-01T12:00:00Z"),
    messaging.WithTemplateUpdateVariableDefaults([]interface{}{}),
    messaging.WithTemplateUpdateVariables([]interface{}{}),
    messaging.WithTemplateUpdateWhatsappCategory("marketing"),
)
```
