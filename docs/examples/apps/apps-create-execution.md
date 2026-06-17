```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/apps"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com")
    client.WithApiKeyAuth("<API_KEY>")
)

service := apps.New(client)

response, error := service.AppsCreateExecution(
    "",
    apps.WithAppsCreateExecutionAsync(false),
    apps.WithAppsCreateExecutionBody(""),
    apps.WithAppsCreateExecutionHeaders(map[string]interface{}{}),
    apps.WithAppsCreateExecutionMethod(""),
    apps.WithAppsCreateExecutionPath(""),
    apps.WithAppsCreateExecutionScheduledAt(""),
)
```
