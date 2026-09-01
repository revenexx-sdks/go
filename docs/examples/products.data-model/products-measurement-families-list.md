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

response, error := service.ProductsMeasurementFamiliesList(
    products_data_model.WithProductsMeasurementFamiliesListLimit(1),
    products_data_model.WithProductsMeasurementFamiliesListOffset(1),
    products_data_model.WithProductsMeasurementFamiliesListOrder("created_at.desc"),
    products_data_model.WithProductsMeasurementFamiliesListId(""),
    products_data_model.WithProductsMeasurementFamiliesListCode("weight"),
    products_data_model.WithProductsMeasurementFamiliesListStandardUnit("kilogram"),
    products_data_model.WithProductsMeasurementFamiliesListUnits("[]"),
    products_data_model.WithProductsMeasurementFamiliesListLabels("{}"),
    products_data_model.WithProductsMeasurementFamiliesListCreatedAt("2026-01-01T12:00:00Z"),
    products_data_model.WithProductsMeasurementFamiliesListUpdatedAt("2026-01-01T12:00:00Z"),
)
```
