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

response, error := service.AppsListTemplates(
    apps.WithAppsListTemplatesRuntimes([]interface{}{}),
    apps.WithAppsListTemplatesUseCases([]interface{}{}),
    apps.WithAppsListTemplatesLimit(0),
    apps.WithAppsListTemplatesOffset(0),
    apps.WithAppsListTemplatesTotal(false),
)
```
