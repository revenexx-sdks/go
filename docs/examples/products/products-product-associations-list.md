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

response, error := service.ProductsProductAssociationsList(
    products.WithProductsProductAssociationsListLimit(1),
    products.WithProductsProductAssociationsListOffset(1),
    products.WithProductsProductAssociationsListOrder("created_at.desc"),
    products.WithProductsProductAssociationsListId(""),
    products.WithProductsProductAssociationsListProductId(""),
    products.WithProductsProductAssociationsListAssociationTypeId(""),
    products.WithProductsProductAssociationsListTargetProductId(""),
    products.WithProductsProductAssociationsListQuantity(9.99),
    products.WithProductsProductAssociationsListPosition(1),
    products.WithProductsProductAssociationsListCreatedAt("2026-01-01T12:00:00Z"),
)
```
