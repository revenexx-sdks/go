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

response, error := service.AppsCreate(
    "",
    "",
    "node-18.0",
    apps.WithAppsCreateCommands("npm install"),
    apps.WithAppsCreateEnabled(true),
    apps.WithAppsCreateEntrypoint("src/main.js"),
    apps.WithAppsCreateEvents([]interface{}{}),
    apps.WithAppsCreateExecute(interface{}{"any"}),
    apps.WithAppsCreateInstallationId(""),
    apps.WithAppsCreateLogging(true),
    apps.WithAppsCreateProviderBranch("main"),
    apps.WithAppsCreateProviderRepositoryId(""),
    apps.WithAppsCreateProviderRootDirectory(""),
    apps.WithAppsCreateProviderSilentMode(true),
    apps.WithAppsCreateSchedule("0 3 * * *"),
    apps.WithAppsCreateScopes([]interface{}{}),
    apps.WithAppsCreateSpecification("s-1vcpu-512mb"),
    apps.WithAppsCreateTimeout(1),
)
```
