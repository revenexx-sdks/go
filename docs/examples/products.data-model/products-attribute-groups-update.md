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

response, error := service.ProductsAttributeGroupsUpdate(
    "",
    products_data_model.WithProductsAttributeGroupsUpdateCode("technical_attributes"),
    products_data_model.WithProductsAttributeGroupsUpdateLabels(map[string]interface{}{
        "de": "Technische Attribute",
        "en": "Technical attributes"
    }),
    products_data_model.WithProductsAttributeGroupsUpdatePosition(1),
)
```
