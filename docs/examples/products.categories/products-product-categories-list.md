```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/products_categories"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := products_categories.New(client)

response, error := service.ProductsProductCategoriesList(
    products_categories.WithProductsProductCategoriesListLimit(1),
    products_categories.WithProductsProductCategoriesListOffset(1),
    products_categories.WithProductsProductCategoriesListOrder("created_at.desc"),
    products_categories.WithProductsProductCategoriesListId(""),
    products_categories.WithProductsProductCategoriesListProductId(""),
    products_categories.WithProductsProductCategoriesListCategoryId(""),
    products_categories.WithProductsProductCategoriesListPosition(1),
    products_categories.WithProductsProductCategoriesListSource("manual"),
    products_categories.WithProductsProductCategoriesListCreatedAt("2026-01-01T12:00:00Z"),
)
```
