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

response, error := service.ProductsMeasurementFamiliesCreate(
    "weight",
    "kilogram",
    products_data_model.WithProductsMeasurementFamiliesCreateLabels(map[string]interface{}{
        "de": "Gewicht",
        "en": "Weight"
    }),
    products_data_model.WithProductsMeasurementFamiliesCreateUnits(map[string]interface{}[
        {
            "code": "kilogram",
            "convert_factor": 1,
            "symbol": "kg"
        },
        {
            "code": "gram",
            "convert_factor": 0.001,
            "symbol": "g"
        }
    ]),
)
```
