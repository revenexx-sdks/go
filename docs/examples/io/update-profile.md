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

response, error := service.UpdateProfile(
    "",
    "",
    "import",
    "",
    "",
    "",
    "",
    io.WithUpdateProfileApplyMode("upsert"),
    io.WithUpdateProfileMapping(map[string]interface{}{}),
    io.WithUpdateProfileMarkets([]interface{}{}),
    io.WithUpdateProfileOptions(map[string]interface{}{}),
)
```
