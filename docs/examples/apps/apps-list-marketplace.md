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

response, error := service.AppsListMarketplace(
    apps.WithAppsListMarketplaceSearch(""),
    apps.WithAppsListMarketplacePerPage(0),
    apps.WithAppsListMarketplacePage(0),
)
```
