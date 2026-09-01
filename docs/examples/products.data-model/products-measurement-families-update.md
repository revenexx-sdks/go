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

response, error := service.ProductsMeasurementFamiliesUpdate(
    "",
    products_data_model.WithProductsMeasurementFamiliesUpdateCode("weight"),
    products_data_model.WithProductsMeasurementFamiliesUpdateLabels(map[string]interface{}{
        "de": "Gewicht",
        "en": "Weight"
    }),
    products_data_model.WithProductsMeasurementFamiliesUpdateStandardUnit("kilogram"),
    products_data_model.WithProductsMeasurementFamiliesUpdateUnits(map[string]interface{}[
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
