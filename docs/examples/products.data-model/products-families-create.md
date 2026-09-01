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

response, error := service.ProductsFamiliesCreate(
    "power_tools",
    products_data_model.WithProductsFamiliesCreateImageAttribute("main_image"),
    products_data_model.WithProductsFamiliesCreateLabelAttribute("name"),
    products_data_model.WithProductsFamiliesCreateLabels(map[string]interface{}{
        "de": "Elektrowerkzeuge",
        "en": "Power tools"
    }),
)
```
