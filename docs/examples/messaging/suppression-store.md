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

response, error := service.SuppressionStore(
    "",
    "",
    "hard_bounce",
    messaging.WithSuppressionStoreExpiresAt("2026-01-01T12:00:00Z"),
    messaging.WithSuppressionStoreNote(""),
    messaging.WithSuppressionStoreScope("all"),
)
```
