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

response, error := service.ProductsAttributeGroupsCreate(
    "technical_attributes",
    products_data_model.WithProductsAttributeGroupsCreateLabels(map[string]interface{}{
        "de": "Technische Attribute",
        "en": "Technical attributes"
    }),
    products_data_model.WithProductsAttributeGroupsCreatePosition(1),
)
```
