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

response, error := service.ProductsAssociationTypesList(
    products_data_model.WithProductsAssociationTypesListLimit(1),
    products_data_model.WithProductsAssociationTypesListOffset(1),
    products_data_model.WithProductsAssociationTypesListOrder("created_at.desc"),
    products_data_model.WithProductsAssociationTypesListId(""),
    products_data_model.WithProductsAssociationTypesListCode("cross_sell"),
    products_data_model.WithProductsAssociationTypesListIsTwoWay(true),
    products_data_model.WithProductsAssociationTypesListIsQuantified(true),
    products_data_model.WithProductsAssociationTypesListLabels("{}"),
    products_data_model.WithProductsAssociationTypesListCreatedAt("2026-01-01T12:00:00Z"),
)
```
