```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/settings"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := settings.New(client)

response, error := service.SettingsGetAppSettings(
    "",
    settings.WithSettingsGetAppSettingsMarket(""),
)
```
