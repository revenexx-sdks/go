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

response, error := service.ProductsAttributeGroupsList(
    products_data_model.WithProductsAttributeGroupsListLimit(1),
    products_data_model.WithProductsAttributeGroupsListOffset(1),
    products_data_model.WithProductsAttributeGroupsListOrder("created_at.desc"),
    products_data_model.WithProductsAttributeGroupsListId(""),
    products_data_model.WithProductsAttributeGroupsListCode("technical_attributes"),
    products_data_model.WithProductsAttributeGroupsListPosition(1),
    products_data_model.WithProductsAttributeGroupsListLabels("{}"),
    products_data_model.WithProductsAttributeGroupsListCreatedAt("2026-01-01T12:00:00Z"),
    products_data_model.WithProductsAttributeGroupsListUpdatedAt("2026-01-01T12:00:00Z"),
)
```
