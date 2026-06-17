```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/storage"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com")
    client.WithApiKeyAuth("<API_KEY>")
)

service := storage.New(client)

response, error := service.AssetUpdate(
    "",
    storage.WithAssetUpdateAltText(""),
    storage.WithAssetUpdateDescription(""),
    storage.WithAssetUpdateDisplayName(""),
    storage.WithAssetUpdateFolderId(""),
    storage.WithAssetUpdateName(""),
    storage.WithAssetUpdateTags([]interface{}{}),
    storage.WithAssetUpdateVisibility(""),
)
```
