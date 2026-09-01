```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/products"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := products.New(client)

response, error := service.ProductsList(
    products.WithProductsListLimit(1),
    products.WithProductsListOffset(1),
    products.WithProductsListOrder("created_at.desc"),
    products.WithProductsListId(""),
    products.WithProductsListSku("ACME-4711-BLK"),
    products.WithProductsListKind("simple"),
    products.WithProductsListParentId(""),
    products.WithProductsListFamilyId(""),
    products.WithProductsListFamilyVariantId(""),
    products.WithProductsListEnabled(true),
    products.WithProductsListTaxClass("standard"),
    products.WithProductsListAttributeValues("{}"),
    products.WithProductsListLabel("Akku-Bohrschrauber 18V"),
    products.WithProductsListQuantifiedAssociations("{}"),
    products.WithProductsListCompleteness("{}"),
    products.WithProductsListCreatedAt("2026-01-01T12:00:00Z"),
    products.WithProductsListUpdatedAt("2026-01-01T12:00:00Z"),
    products.WithProductsListDeletedAt("2026-01-01T12:00:00Z"),
)
```
