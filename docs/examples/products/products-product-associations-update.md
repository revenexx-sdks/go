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

response, error := service.ProductsProductAssociationsUpdate(
    "",
    products.WithProductsProductAssociationsUpdateAssociationTypeId(""),
    products.WithProductsProductAssociationsUpdatePosition(0),
    products.WithProductsProductAssociationsUpdateProductId(""),
    products.WithProductsProductAssociationsUpdateQuantity(0),
    products.WithProductsProductAssociationsUpdateTargetProductId(""),
)
```
