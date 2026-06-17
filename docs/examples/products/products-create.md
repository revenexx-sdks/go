```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/products"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com")
    client.WithApiKeyAuth("<API_KEY>")
)

service := products.New(client)

response, error := service.ProductsCreate(
    "",
    products.WithProductsCreateAttributeValues(map[string]interface{}{}),
    products.WithProductsCreateCompleteness(map[string]interface{}{}),
    products.WithProductsCreateDeletedAt(""),
    products.WithProductsCreateEnabled(false),
    products.WithProductsCreateFamilyId(""),
    products.WithProductsCreateFamilyVariantId(""),
    products.WithProductsCreateKind(""),
    products.WithProductsCreateParentId(""),
    products.WithProductsCreateQuantifiedAssociations(map[string]interface{}{}),
    products.WithProductsCreateTaxClass(""),
)
```
