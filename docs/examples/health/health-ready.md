```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/health"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := health.New(client)

response, error := service.HealthReady()
```
