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

response, error := service.ProductsFamilyAttributesCreate(
    "",
    "",
    products_data_model.WithProductsFamilyAttributesCreateIsRequired(true),
    products_data_model.WithProductsFamilyAttributesCreatePosition(1),
    products_data_model.WithProductsFamilyAttributesCreateRequiredChannels(map[string]interface{}[
        "shop",
        "b2b"
    ]),
)
```
