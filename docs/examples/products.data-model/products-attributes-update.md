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

response, error := service.ProductsAttributesUpdate(
    "",
    products_data_model.WithProductsAttributesUpdateCode("net_weight"),
    products_data_model.WithProductsAttributesUpdateConfig(map[string]interface{}{
        "reference_entity": "brand"
    }),
    products_data_model.WithProductsAttributesUpdateEntityRef("brand"),
    products_data_model.WithProductsAttributesUpdateEntityType("product"),
    products_data_model.WithProductsAttributesUpdateGroupId(""),
    products_data_model.WithProductsAttributesUpdateIsFilterable(true),
    products_data_model.WithProductsAttributesUpdateIsUnique(true),
    products_data_model.WithProductsAttributesUpdateLabels(map[string]interface{}{
        "de": "Nettogewicht",
        "en": "Net weight"
    }),
    products_data_model.WithProductsAttributesUpdateLocalizable(true),
    products_data_model.WithProductsAttributesUpdatePosition(1),
    products_data_model.WithProductsAttributesUpdateScopable(true),
    products_data_model.WithProductsAttributesUpdateType("select"),
    products_data_model.WithProductsAttributesUpdateUsableInGrid(true),
    products_data_model.WithProductsAttributesUpdateValidation(map[string]interface{}{
        "max_length": 64,
        "min_length": 3
    }),
)
```
