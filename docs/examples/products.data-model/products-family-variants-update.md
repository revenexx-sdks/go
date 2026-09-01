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

response, error := service.ProductsFamilyVariantsUpdate(
    "",
    products_data_model.WithProductsFamilyVariantsUpdateAxes(map[string]interface{}[
        "colour",
        "size"
    ]),
    products_data_model.WithProductsFamilyVariantsUpdateCode("clothing_by_colour_size"),
    products_data_model.WithProductsFamilyVariantsUpdateFamilyId(""),
    products_data_model.WithProductsFamilyVariantsUpdateLabels(map[string]interface{}{
        "de": "Nach Farbe und Gr\u00f6\u00dfe",
        "en": "By colour and size"
    }),
)
```
