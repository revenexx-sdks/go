```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/io"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := io.New(client)

response, error := service.CreateImport(
    "",
    "",
    "",
    "",
    io.WithCreateImportFormat("csv"),
    io.WithCreateImportKeys([]interface{}{}),
    io.WithCreateImportMaxRejects(1),
    io.WithCreateImportMode("upsert"),
    io.WithCreateImportProfileId(""),
    io.WithCreateImportTarget("live"),
)
```
