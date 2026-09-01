```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/products_references"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := products_references.New(client)

response, error := service.ProductsReferenceEntitiesUpdate(
    "",
    products_references.WithProductsReferenceEntitiesUpdateCode("brand"),
    products_references.WithProductsReferenceEntitiesUpdateImage("reference-entities/brand.svg"),
    products_references.WithProductsReferenceEntitiesUpdateLabels(map[string]interface{}{
        "de": "Marke",
        "en": "Brand"
    }),
)
```
