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

response, error := service.ProductsAttributeOptionsCreate(
    "",
    "stainless_steel",
    products_data_model.WithProductsAttributeOptionsCreateLabels(map[string]interface{}{
        "de": "Edelstahl",
        "en": "Stainless steel"
    }),
    products_data_model.WithProductsAttributeOptionsCreatePosition(1),
    products_data_model.WithProductsAttributeOptionsCreateSwatch(map[string]interface{}{
        "hex": "#c0c0c0"
    }),
)
```
