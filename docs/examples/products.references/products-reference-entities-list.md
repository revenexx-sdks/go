```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/products_references"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := products_references.New(client)

response, error := service.ProductsReferenceEntitiesList(
    products_references.WithProductsReferenceEntitiesListLimit(1),
    products_references.WithProductsReferenceEntitiesListOffset(1),
    products_references.WithProductsReferenceEntitiesListOrder("created_at.desc"),
    products_references.WithProductsReferenceEntitiesListId(""),
    products_references.WithProductsReferenceEntitiesListCode("brand"),
    products_references.WithProductsReferenceEntitiesListLabels("{}"),
    products_references.WithProductsReferenceEntitiesListImage("reference-entities/brand.svg"),
    products_references.WithProductsReferenceEntitiesListCreatedAt("2026-01-01T12:00:00Z"),
    products_references.WithProductsReferenceEntitiesListUpdatedAt("2026-01-01T12:00:00Z"),
)
```
