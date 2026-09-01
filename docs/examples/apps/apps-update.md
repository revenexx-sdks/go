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

response, error := service.AppsUpdate(
    "",
    "",
    apps.WithAppsUpdateCommands("npm install"),
    apps.WithAppsUpdateEnabled(true),
    apps.WithAppsUpdateEntrypoint("src/main.js"),
    apps.WithAppsUpdateEvents([]interface{}{}),
    apps.WithAppsUpdateExecute(interface{}{"any"}),
    apps.WithAppsUpdateInstallationId(""),
    apps.WithAppsUpdateLogging(true),
    apps.WithAppsUpdateProviderBranch("main"),
    apps.WithAppsUpdateProviderRepositoryId(""),
    apps.WithAppsUpdateProviderRootDirectory(""),
    apps.WithAppsUpdateProviderSilentMode(true),
    apps.WithAppsUpdateRuntime("node-18.0"),
    apps.WithAppsUpdateSchedule("0 3 * * *"),
    apps.WithAppsUpdateScopes([]interface{}{}),
    apps.WithAppsUpdateSpecification("s-1vcpu-512mb"),
    apps.WithAppsUpdateTimeout(1),
)
```
