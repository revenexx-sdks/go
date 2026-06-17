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

response, error := service.AppsUpdate(
    "",
    "",
    apps.WithAppsUpdateCommands(""),
    apps.WithAppsUpdateEnabled(false),
    apps.WithAppsUpdateEntrypoint(""),
    apps.WithAppsUpdateEvents([]interface{}{}),
    apps.WithAppsUpdateExecute([]interface{}{}),
    apps.WithAppsUpdateInstallationId(""),
    apps.WithAppsUpdateLogging(false),
    apps.WithAppsUpdateProviderBranch(""),
    apps.WithAppsUpdateProviderRepositoryId(""),
    apps.WithAppsUpdateProviderRootDirectory(""),
    apps.WithAppsUpdateProviderSilentMode(false),
    apps.WithAppsUpdateRuntime(""),
    apps.WithAppsUpdateSchedule(""),
    apps.WithAppsUpdateScopes([]interface{}{}),
    apps.WithAppsUpdateSpecification(""),
    apps.WithAppsUpdateTimeout(0),
)
```
