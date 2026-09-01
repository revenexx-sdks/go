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

response, error := service.ProductsAssetsList(
    products_assets.WithProductsAssetsListLimit(1),
    products_assets.WithProductsAssetsListOffset(1),
    products_assets.WithProductsAssetsListOrder("created_at.desc"),
    products_assets.WithProductsAssetsListId(""),
    products_assets.WithProductsAssetsListAssetFamilyId(""),
    products_assets.WithProductsAssetsListCode("acme-4711-blk_packshot_1"),
    products_assets.WithProductsAssetsListSource("storage"),
    products_assets.WithProductsAssetsListStorageAssetId("ast_01J8ZQ0000000000000000"),
    products_assets.WithProductsAssetsListDeliveryPath("packshots/acme-4711-blk_1.jpg"),
    products_assets.WithProductsAssetsListExternalUrl("https://cdn.example.com/packshots/acme-4711-blk_1.jpg"),
    products_assets.WithProductsAssetsListAttributeValues("{}"),
    products_assets.WithProductsAssetsListCreatedAt("2026-01-01T12:00:00Z"),
    products_assets.WithProductsAssetsListUpdatedAt("2026-01-01T12:00:00Z"),
)
```
