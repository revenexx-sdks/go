```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/sites"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com")
    client.WithApiKeyAuth("<API_KEY>")
)

service := sites.New(client)

response, error := service.SitesCreate(
    "",
    "",
    "",
    "",
    sites.WithSitesCreateAdapter(""),
    sites.WithSitesCreateBuildCommand(""),
    sites.WithSitesCreateEnabled(false),
    sites.WithSitesCreateFallbackFile(""),
    sites.WithSitesCreateInstallCommand(""),
    sites.WithSitesCreateInstallationId(""),
    sites.WithSitesCreateLogging(false),
    sites.WithSitesCreateOutputDirectory(""),
    sites.WithSitesCreateProviderBranch(""),
    sites.WithSitesCreateProviderRepositoryId(""),
    sites.WithSitesCreateProviderRootDirectory(""),
    sites.WithSitesCreateProviderSilentMode(false),
    sites.WithSitesCreateSpecification(""),
    sites.WithSitesCreateTimeout(0),
)
```
