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

response, error := service.ProductsGrid(
    products.WithProductsGridLimit(1),
    products.WithProductsGridOffset(1),
    products.WithProductsGridOrder("created_at.desc"),
    products.WithProductsGridQ("cordless drill"),
    products.WithProductsGridKind("simple"),
    products.WithProductsGridEnabled(true),
    products.WithProductsGridFamilyId(""),
)
```
