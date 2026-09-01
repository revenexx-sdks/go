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

response, error := service.ProductsAssetsUpdate(
    "",
    products_assets.WithProductsAssetsUpdateAssetFamilyId(""),
    products_assets.WithProductsAssetsUpdateAttributeValues(map[string]interface{}{
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
    products_assets.WithProductsAssetsUpdateCode("acme-4711-blk_packshot_1"),
    products_assets.WithProductsAssetsUpdateDeliveryPath("packshots/acme-4711-blk_1.jpg"),
    products_assets.WithProductsAssetsUpdateExternalUrl("https://cdn.example.com/packshots/acme-4711-blk_1.jpg"),
    products_assets.WithProductsAssetsUpdateSource("storage"),
    products_assets.WithProductsAssetsUpdateStorageAssetId("ast_01J8ZQ0000000000000000"),
)
```
