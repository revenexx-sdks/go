```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/sites"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := sites.New(client)

response, error := service.SitesUpdate(
    "",
    "analog",
    "",
    sites.WithSitesUpdateAdapter("static"),
    sites.WithSitesUpdateBuildCommand("npm run build"),
    sites.WithSitesUpdateBuildRuntime("node-18.0"),
    sites.WithSitesUpdateEnabled(true),
    sites.WithSitesUpdateFallbackFile("index.html"),
    sites.WithSitesUpdateInstallCommand("npm install"),
    sites.WithSitesUpdateInstallationId(""),
    sites.WithSitesUpdateLogging(true),
    sites.WithSitesUpdateOutputDirectory(""),
    sites.WithSitesUpdateProviderBranch("main"),
    sites.WithSitesUpdateProviderRepositoryId(""),
    sites.WithSitesUpdateProviderRootDirectory(""),
    sites.WithSitesUpdateProviderSilentMode(true),
    sites.WithSitesUpdateSpecification("s-1vcpu-512mb"),
    sites.WithSitesUpdateTimeout(1),
)
```
