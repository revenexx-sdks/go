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

response, error := service.OrderlistsList(
    orderlists.WithOrderlistsListOwnerId(""),
    orderlists.WithOrderlistsListOrganizationId(""),
    orderlists.WithOrderlistsListKind("shopping"),
    orderlists.WithOrderlistsListLimit(50),
    orderlists.WithOrderlistsListOffset(0),
    orderlists.WithOrderlistsListOrder("created_at.desc"),
)
```
