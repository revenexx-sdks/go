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

response, error := service.ProductsAttributesUpdate(
    "",
    products.WithProductsAttributesUpdateCode(""),
    products.WithProductsAttributesUpdateConfig(map[string]interface{}{}),
    products.WithProductsAttributesUpdateEntityRef(""),
    products.WithProductsAttributesUpdateEntityType(""),
    products.WithProductsAttributesUpdateGroupId(""),
    products.WithProductsAttributesUpdateIsFilterable(false),
    products.WithProductsAttributesUpdateIsUnique(false),
    products.WithProductsAttributesUpdateLabels(map[string]interface{}{}),
    products.WithProductsAttributesUpdateLocalizable(false),
    products.WithProductsAttributesUpdatePosition(0),
    products.WithProductsAttributesUpdateScopable(false),
    products.WithProductsAttributesUpdateType(""),
    products.WithProductsAttributesUpdateUsableInGrid(false),
    products.WithProductsAttributesUpdateValidation(map[string]interface{}{}),
)
```
