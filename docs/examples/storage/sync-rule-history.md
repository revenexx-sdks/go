```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/storage"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := storage.New(client)

response, error := service.SyncRuleHistory(
    storage.WithSyncRuleHistoryRuleId(""),
    storage.WithSyncRuleHistoryFrom("2026-01-01T12:00:00Z"),
    storage.WithSyncRuleHistoryTo("2026-01-01T12:00:00Z"),
)
```
