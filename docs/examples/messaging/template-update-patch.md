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

response, error := service.TemplateUpdatePatch(
    "",
    messaging.WithTemplateUpdatePatchBodyHtml(""),
    messaging.WithTemplateUpdatePatchBodyText(""),
    messaging.WithTemplateUpdatePatchContentSid(""),
    messaging.WithTemplateUpdatePatchDesign([]interface{}{}),
    messaging.WithTemplateUpdatePatchEnabled(true),
    messaging.WithTemplateUpdatePatchLayoutId(""),
    messaging.WithTemplateUpdatePatchMarkets([]interface{}{}),
    messaging.WithTemplateUpdatePatchMessageClass("transactional"),
    messaging.WithTemplateUpdatePatchSubject(""),
    messaging.WithTemplateUpdatePatchTestMode(true),
    messaging.WithTemplateUpdatePatchTitle(""),
    messaging.WithTemplateUpdatePatchValidFrom("2026-01-01T12:00:00Z"),
    messaging.WithTemplateUpdatePatchValidUntil("2026-01-01T12:00:00Z"),
    messaging.WithTemplateUpdatePatchVariableDefaults([]interface{}{}),
    messaging.WithTemplateUpdatePatchVariables([]interface{}{}),
    messaging.WithTemplateUpdatePatchWhatsappCategory("marketing"),
)
```
