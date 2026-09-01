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

response, error := service.ProductsFamilyAttributesList(
    products_data_model.WithProductsFamilyAttributesListLimit(1),
    products_data_model.WithProductsFamilyAttributesListOffset(1),
    products_data_model.WithProductsFamilyAttributesListOrder("created_at.desc"),
    products_data_model.WithProductsFamilyAttributesListId(""),
    products_data_model.WithProductsFamilyAttributesListFamilyId(""),
    products_data_model.WithProductsFamilyAttributesListAttributeId(""),
    products_data_model.WithProductsFamilyAttributesListPosition(1),
    products_data_model.WithProductsFamilyAttributesListIsRequired(true),
    products_data_model.WithProductsFamilyAttributesListRequiredChannels("[]"),
    products_data_model.WithProductsFamilyAttributesListCreatedAt("2026-01-01T12:00:00Z"),
)
```
