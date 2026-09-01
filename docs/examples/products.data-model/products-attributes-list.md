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

response, error := service.ProductsAttributesList(
    products_data_model.WithProductsAttributesListLimit(1),
    products_data_model.WithProductsAttributesListOffset(1),
    products_data_model.WithProductsAttributesListOrder("created_at.desc"),
    products_data_model.WithProductsAttributesListId(""),
    products_data_model.WithProductsAttributesListCode("net_weight"),
    products_data_model.WithProductsAttributesListEntityType("product"),
    products_data_model.WithProductsAttributesListEntityRef("brand"),
    products_data_model.WithProductsAttributesListType("select"),
    products_data_model.WithProductsAttributesListGroupId(""),
    products_data_model.WithProductsAttributesListLocalizable(true),
    products_data_model.WithProductsAttributesListScopable(true),
    products_data_model.WithProductsAttributesListIsUnique(true),
    products_data_model.WithProductsAttributesListIsFilterable(true),
    products_data_model.WithProductsAttributesListUsableInGrid(true),
    products_data_model.WithProductsAttributesListValidation("{}"),
    products_data_model.WithProductsAttributesListConfig("{}"),
    products_data_model.WithProductsAttributesListLabels("{}"),
    products_data_model.WithProductsAttributesListPosition(1),
    products_data_model.WithProductsAttributesListCreatedAt("2026-01-01T12:00:00Z"),
    products_data_model.WithProductsAttributesListUpdatedAt("2026-01-01T12:00:00Z"),
)
```
