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

response, error := service.ProductsUpdate(
    "",
    products.WithProductsUpdateAttributeValues(map[string]interface{}{}),
    products.WithProductsUpdateCompleteness(map[string]interface{}{}),
    products.WithProductsUpdateDeletedAt(""),
    products.WithProductsUpdateEnabled(false),
    products.WithProductsUpdateFamilyId(""),
    products.WithProductsUpdateFamilyVariantId(""),
    products.WithProductsUpdateKind(""),
    products.WithProductsUpdateParentId(""),
    products.WithProductsUpdateQuantifiedAssociations(map[string]interface{}{}),
    products.WithProductsUpdateSku(""),
    products.WithProductsUpdateTaxClass(""),
)
```
