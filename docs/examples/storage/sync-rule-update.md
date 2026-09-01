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

response, error := service.SyncRuleUpdate(
    "",
    storage.WithSyncRuleUpdateEnabled(true),
    storage.WithSyncRuleUpdateOptions([]interface{}{}),
    storage.WithSyncRuleUpdateSchedule("0 3 * * *"),
    storage.WithSyncRuleUpdateSftpAccountId(""),
    storage.WithSyncRuleUpdateSourcePath("/uploads"),
    storage.WithSyncRuleUpdateTargetFolderId(""),
)
```
