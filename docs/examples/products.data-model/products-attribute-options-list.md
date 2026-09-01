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

response, error := service.ProductsAttributeOptionsList(
    products_data_model.WithProductsAttributeOptionsListLimit(1),
    products_data_model.WithProductsAttributeOptionsListOffset(1),
    products_data_model.WithProductsAttributeOptionsListOrder("created_at.desc"),
    products_data_model.WithProductsAttributeOptionsListId(""),
    products_data_model.WithProductsAttributeOptionsListAttributeId(""),
    products_data_model.WithProductsAttributeOptionsListCode("stainless_steel"),
    products_data_model.WithProductsAttributeOptionsListPosition(1),
    products_data_model.WithProductsAttributeOptionsListSwatch("{}"),
    products_data_model.WithProductsAttributeOptionsListLabels("{}"),
    products_data_model.WithProductsAttributeOptionsListCreatedAt("2026-01-01T12:00:00Z"),
)
```
