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

response, error := service.ProductsAttributesCreate(
    "net_weight",
    "select",
    products_data_model.WithProductsAttributesCreateConfig(map[string]interface{}{
        "reference_entity": "brand"
    }),
    products_data_model.WithProductsAttributesCreateEntityRef("brand"),
    products_data_model.WithProductsAttributesCreateEntityType("product"),
    products_data_model.WithProductsAttributesCreateGroupId(""),
    products_data_model.WithProductsAttributesCreateIsFilterable(true),
    products_data_model.WithProductsAttributesCreateIsUnique(true),
    products_data_model.WithProductsAttributesCreateLabels(map[string]interface{}{
        "de": "Nettogewicht",
        "en": "Net weight"
    }),
    products_data_model.WithProductsAttributesCreateLocalizable(true),
    products_data_model.WithProductsAttributesCreatePosition(1),
    products_data_model.WithProductsAttributesCreateScopable(true),
    products_data_model.WithProductsAttributesCreateUsableInGrid(true),
    products_data_model.WithProductsAttributesCreateValidation(map[string]interface{}{
        "max_length": 64,
        "min_length": 3
    }),
)
```
