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

response, error := service.ProductsProductAssociationsCreate(
    "",
    "",
    "",
    products.WithProductsProductAssociationsCreatePosition(0),
    products.WithProductsProductAssociationsCreateQuantity(0),
)
```
