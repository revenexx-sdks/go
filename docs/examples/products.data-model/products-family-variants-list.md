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

response, error := service.ProductsFamilyVariantsList(
    products_data_model.WithProductsFamilyVariantsListLimit(1),
    products_data_model.WithProductsFamilyVariantsListOffset(1),
    products_data_model.WithProductsFamilyVariantsListOrder("created_at.desc"),
    products_data_model.WithProductsFamilyVariantsListId(""),
    products_data_model.WithProductsFamilyVariantsListFamilyId(""),
    products_data_model.WithProductsFamilyVariantsListCode("clothing_by_colour_size"),
    products_data_model.WithProductsFamilyVariantsListLabels("{}"),
    products_data_model.WithProductsFamilyVariantsListAxes("[]"),
    products_data_model.WithProductsFamilyVariantsListCreatedAt("2026-01-01T12:00:00Z"),
    products_data_model.WithProductsFamilyVariantsListUpdatedAt("2026-01-01T12:00:00Z"),
)
```
