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

response, error := service.ProductsFamiliesList(
    products_data_model.WithProductsFamiliesListLimit(1),
    products_data_model.WithProductsFamiliesListOffset(1),
    products_data_model.WithProductsFamiliesListOrder("created_at.desc"),
    products_data_model.WithProductsFamiliesListId(""),
    products_data_model.WithProductsFamiliesListCode("power_tools"),
    products_data_model.WithProductsFamiliesListLabelAttribute("name"),
    products_data_model.WithProductsFamiliesListImageAttribute("main_image"),
    products_data_model.WithProductsFamiliesListLabels("{}"),
    products_data_model.WithProductsFamiliesListCreatedAt("2026-01-01T12:00:00Z"),
    products_data_model.WithProductsFamiliesListUpdatedAt("2026-01-01T12:00:00Z"),
)
```
