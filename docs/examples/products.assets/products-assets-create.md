```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/products_assets"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := products_assets.New(client)

response, error := service.ProductsAssetsCreate(
    "",
    "acme-4711-blk_packshot_1",
    products_assets.WithProductsAssetsCreateAttributeValues(map[string]interface{}{
        "common": {
            "copyright": "\u00a9 Acme Tools",
            "expires_on": "2028-12-31"
        },
        "locale_specific": {
            "de_DE": {
                "alt_text": "Akku-Bohrschrauber, freigestellt"
            }
        }
    }),
    products_assets.WithProductsAssetsCreateDeliveryPath("packshots/acme-4711-blk_1.jpg"),
    products_assets.WithProductsAssetsCreateExternalUrl("https://cdn.example.com/packshots/acme-4711-blk_1.jpg"),
    products_assets.WithProductsAssetsCreateSource("storage"),
    products_assets.WithProductsAssetsCreateStorageAssetId("ast_01J8ZQ0000000000000000"),
)
```
