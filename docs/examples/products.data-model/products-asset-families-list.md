```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/products_data_model"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := products_data_model.New(client)

response, error := service.ProductsAssetFamiliesList(
    products_data_model.WithProductsAssetFamiliesListLimit(1),
    products_data_model.WithProductsAssetFamiliesListOffset(1),
    products_data_model.WithProductsAssetFamiliesListOrder("created_at.desc"),
    products_data_model.WithProductsAssetFamiliesListId(""),
    products_data_model.WithProductsAssetFamiliesListCode("packshots"),
    products_data_model.WithProductsAssetFamiliesListLabels("{}"),
    products_data_model.WithProductsAssetFamiliesListNamingConvention("{}"),
    products_data_model.WithProductsAssetFamiliesListCreatedAt("2026-01-01T12:00:00Z"),
    products_data_model.WithProductsAssetFamiliesListUpdatedAt("2026-01-01T12:00:00Z"),
)
```
