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

response, error := service.AppsListTemplates(
    apps.WithAppsListTemplatesRuntimes([]interface{}{}),
    apps.WithAppsListTemplatesUseCases([]interface{}{}),
    apps.WithAppsListTemplatesLimit(1),
    apps.WithAppsListTemplatesOffset(1),
    apps.WithAppsListTemplatesTotal(true),
)
```
