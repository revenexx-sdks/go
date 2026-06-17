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

response, error := service.ProductsCategoriesUpdate(
    "",
    products.WithProductsCategoriesUpdateCode(""),
    products.WithProductsCategoriesUpdateLabels(map[string]interface{}{}),
    products.WithProductsCategoriesUpdateParentId(""),
    products.WithProductsCategoriesUpdatePath(""),
    products.WithProductsCategoriesUpdatePosition(0),
    products.WithProductsCategoriesUpdateValues(map[string]interface{}{}),
)
```
