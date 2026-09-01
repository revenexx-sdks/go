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

response, error := service.SitesCreate(
    "node-18.0",
    "analog",
    "",
    "",
    sites.WithSitesCreateAdapter("static"),
    sites.WithSitesCreateBuildCommand("npm run build"),
    sites.WithSitesCreateEnabled(true),
    sites.WithSitesCreateFallbackFile("index.html"),
    sites.WithSitesCreateInstallCommand("npm install"),
    sites.WithSitesCreateInstallationId(""),
    sites.WithSitesCreateLogging(true),
    sites.WithSitesCreateOutputDirectory(""),
    sites.WithSitesCreateProviderBranch("main"),
    sites.WithSitesCreateProviderRepositoryId(""),
    sites.WithSitesCreateProviderRootDirectory(""),
    sites.WithSitesCreateProviderSilentMode(true),
    sites.WithSitesCreateSpecification("s-1vcpu-512mb"),
    sites.WithSitesCreateTimeout(1),
)
```
