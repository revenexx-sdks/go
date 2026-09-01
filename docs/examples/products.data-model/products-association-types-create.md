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

response, error := service.ProductsAssociationTypesCreate(
    "cross_sell",
    products_data_model.WithProductsAssociationTypesCreateIsQuantified(true),
    products_data_model.WithProductsAssociationTypesCreateIsTwoWay(true),
    products_data_model.WithProductsAssociationTypesCreateLabels(map[string]interface{}{
        "de": "Querverkauf",
        "en": "Cross-sell"
    }),
)
```
