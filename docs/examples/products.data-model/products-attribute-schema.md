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

response, error := service.ProductsAttributeSchema(
    products_data_model.WithProductsAttributeSchemaFamilyId(""),
    products_data_model.WithProductsAttributeSchemaFamilyCode(""),
    products_data_model.WithProductsAttributeSchemaEntityType("product"),
    products_data_model.WithProductsAttributeSchemaEntityRef("brand"),
    products_data_model.WithProductsAttributeSchemaLocale("de_DE"),
    products_data_model.WithProductsAttributeSchemaChannel("b2b"),
    products_data_model.WithProductsAttributeSchemaKind("simple"),
)
```
