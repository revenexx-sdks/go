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

response, error := service.StatsIndex(
    messaging.WithStatsIndexDays(1),
    messaging.WithStatsIndexFrom("2026-01-01"),
    messaging.WithStatsIndexTo("2026-01-01"),
)
```
