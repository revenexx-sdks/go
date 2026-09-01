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

response, error := service.ProductsAssetFamiliesCreate(
    "packshots",
    products_data_model.WithProductsAssetFamiliesCreateLabels(map[string]interface{}{
        "de": "Packshots",
        "en": "Packshots"
    }),
    products_data_model.WithProductsAssetFamiliesCreateNamingConvention(map[string]interface{}{
        "allowed_extensions": [
            "jpg",
            "png"
        ],
        "pattern": "{sku}_{index}",
        "source": "sku"
    }),
)
```
