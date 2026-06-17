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

response, error := service.SitesUpdate(
    "",
    "",
    "",
    sites.WithSitesUpdateAdapter(""),
    sites.WithSitesUpdateBuildCommand(""),
    sites.WithSitesUpdateBuildRuntime(""),
    sites.WithSitesUpdateEnabled(false),
    sites.WithSitesUpdateFallbackFile(""),
    sites.WithSitesUpdateInstallCommand(""),
    sites.WithSitesUpdateInstallationId(""),
    sites.WithSitesUpdateLogging(false),
    sites.WithSitesUpdateOutputDirectory(""),
    sites.WithSitesUpdateProviderBranch(""),
    sites.WithSitesUpdateProviderRepositoryId(""),
    sites.WithSitesUpdateProviderRootDirectory(""),
    sites.WithSitesUpdateProviderSilentMode(false),
    sites.WithSitesUpdateSpecification(""),
    sites.WithSitesUpdateTimeout(0),
)
```
