```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/storage"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := storage.New(client)

response, error := service.AssetStore(
    file.NewInputFile("/path/to/file.png", "file.png"),
    storage.WithAssetStoreAltText(""),
    storage.WithAssetStoreDescription(""),
    storage.WithAssetStoreDisplayName(""),
    storage.WithAssetStoreFolderId(""),
    storage.WithAssetStoreKeepArchive(true),
    storage.WithAssetStoreTags([]interface{}{}),
    storage.WithAssetStoreUnpack(true),
    storage.WithAssetStoreVisibility("public"),
)
```
