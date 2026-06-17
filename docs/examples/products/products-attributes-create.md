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

response, error := service.ProductsAttributesCreate(
    "",
    "",
    products.WithProductsAttributesCreateConfig(map[string]interface{}{}),
    products.WithProductsAttributesCreateEntityRef(""),
    products.WithProductsAttributesCreateEntityType(""),
    products.WithProductsAttributesCreateGroupId(""),
    products.WithProductsAttributesCreateIsFilterable(false),
    products.WithProductsAttributesCreateIsUnique(false),
    products.WithProductsAttributesCreateLabels(map[string]interface{}{}),
    products.WithProductsAttributesCreateLocalizable(false),
    products.WithProductsAttributesCreatePosition(0),
    products.WithProductsAttributesCreateScopable(false),
    products.WithProductsAttributesCreateUsableInGrid(false),
    products.WithProductsAttributesCreateValidation(map[string]interface{}{}),
)
```
