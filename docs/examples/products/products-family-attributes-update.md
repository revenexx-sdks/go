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

response, error := service.ProductsFamilyAttributesUpdate(
    "",
    products.WithProductsFamilyAttributesUpdateAttributeId(""),
    products.WithProductsFamilyAttributesUpdateFamilyId(""),
    products.WithProductsFamilyAttributesUpdateIsRequired(false),
    products.WithProductsFamilyAttributesUpdatePosition(0),
    products.WithProductsFamilyAttributesUpdateRequiredChannels(map[string]interface{}{}),
)
```
