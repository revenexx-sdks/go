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

response, error := service.ProductsFamiliesUpdate(
    "",
    products_data_model.WithProductsFamiliesUpdateCode("power_tools"),
    products_data_model.WithProductsFamiliesUpdateImageAttribute("main_image"),
    products_data_model.WithProductsFamiliesUpdateLabelAttribute("name"),
    products_data_model.WithProductsFamiliesUpdateLabels(map[string]interface{}{
        "de": "Elektrowerkzeuge",
        "en": "Power tools"
    }),
)
```
