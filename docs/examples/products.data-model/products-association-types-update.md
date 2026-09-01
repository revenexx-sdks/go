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

response, error := service.ProductsAssociationTypesUpdate(
    "",
    products_data_model.WithProductsAssociationTypesUpdateCode("cross_sell"),
    products_data_model.WithProductsAssociationTypesUpdateIsQuantified(true),
    products_data_model.WithProductsAssociationTypesUpdateIsTwoWay(true),
    products_data_model.WithProductsAssociationTypesUpdateLabels(map[string]interface{}{
        "de": "Querverkauf",
        "en": "Cross-sell"
    }),
)
```
