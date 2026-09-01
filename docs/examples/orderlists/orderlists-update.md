```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/orderlists"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := orderlists.New(client)

response, error := service.OrderlistsUpdate(
    "",
    orderlists.WithOrderlistsUpdateKind("shopping"),
    orderlists.WithOrderlistsUpdateMetadata(map[string]interface{}{
        "department": "facility",
        "erp_reference": "REQ-2026-0042"
    }),
    orderlists.WithOrderlistsUpdateName("Weekly office supplies"),
    orderlists.WithOrderlistsUpdateShared(true),
)
```
