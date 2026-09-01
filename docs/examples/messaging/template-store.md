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

response, error := service.TemplateStore(
    "",
    "",
    messaging.WithTemplateStoreBodyHtml(""),
    messaging.WithTemplateStoreBodyText(""),
    messaging.WithTemplateStoreContentSid(""),
    messaging.WithTemplateStoreDesign([]interface{}{}),
    messaging.WithTemplateStoreEnabled(true),
    messaging.WithTemplateStoreLayoutId(""),
    messaging.WithTemplateStoreLocale(""),
    messaging.WithTemplateStoreMarkets([]interface{}{}),
    messaging.WithTemplateStoreMessageClass("transactional"),
    messaging.WithTemplateStoreSubject(""),
    messaging.WithTemplateStoreTestMode(true),
    messaging.WithTemplateStoreTitle(""),
    messaging.WithTemplateStoreValidFrom("2026-01-01T12:00:00Z"),
    messaging.WithTemplateStoreValidUntil("2026-01-01T12:00:00Z"),
    messaging.WithTemplateStoreVariableDefaults([]interface{}{}),
    messaging.WithTemplateStoreVariables([]interface{}{}),
    messaging.WithTemplateStoreWhatsappCategory("marketing"),
)
```
