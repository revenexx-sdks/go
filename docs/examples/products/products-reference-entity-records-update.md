```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/products"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com")
    client.WithApiKeyAuth("<API_KEY>")
)

service := products.New(client)

response, error := service.ProductsReferenceEntityRecordsUpdate(
    "",
    products.WithProductsReferenceEntityRecordsUpdateAttributeValues(map[string]interface{}{}),
    products.WithProductsReferenceEntityRecordsUpdateCode(""),
    products.WithProductsReferenceEntityRecordsUpdateLabels(map[string]interface{}{}),
    products.WithProductsReferenceEntityRecordsUpdateReferenceEntityId(""),
)
```
