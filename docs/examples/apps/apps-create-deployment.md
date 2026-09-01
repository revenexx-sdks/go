```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/apps"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := apps.New(client)

response, error := service.AppsCreateDeployment(
    "",
    true,
    file.NewInputFile("/path/to/file.png", "file.png"),
    apps.WithAppsCreateDeploymentCommands(""),
    apps.WithAppsCreateDeploymentEntrypoint(""),
)
```
