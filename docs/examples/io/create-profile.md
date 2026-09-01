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

response, error := service.CreateProfile(
    "",
    "import",
    "",
    "",
    "",
    "",
    io.WithCreateProfileApplyMode("upsert"),
    io.WithCreateProfileMapping(map[string]interface{}{}),
    io.WithCreateProfileMarkets([]interface{}{}),
    io.WithCreateProfileOptions(map[string]interface{}{}),
)
```
