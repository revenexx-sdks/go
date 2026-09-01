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

response, error := service.ProductsAttributeOptionsUpdate(
    "",
    products_data_model.WithProductsAttributeOptionsUpdateAttributeId(""),
    products_data_model.WithProductsAttributeOptionsUpdateCode("stainless_steel"),
    products_data_model.WithProductsAttributeOptionsUpdateLabels(map[string]interface{}{
        "de": "Edelstahl",
        "en": "Stainless steel"
    }),
    products_data_model.WithProductsAttributeOptionsUpdatePosition(1),
    products_data_model.WithProductsAttributeOptionsUpdateSwatch(map[string]interface{}{
        "hex": "#c0c0c0"
    }),
)
```
