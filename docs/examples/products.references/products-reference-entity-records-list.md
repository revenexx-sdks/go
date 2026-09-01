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

response, error := service.ProductsReferenceEntityRecordsList(
    products_references.WithProductsReferenceEntityRecordsListLimit(1),
    products_references.WithProductsReferenceEntityRecordsListOffset(1),
    products_references.WithProductsReferenceEntityRecordsListOrder("created_at.desc"),
    products_references.WithProductsReferenceEntityRecordsListId(""),
    products_references.WithProductsReferenceEntityRecordsListReferenceEntityId(""),
    products_references.WithProductsReferenceEntityRecordsListCode("acme_tools"),
    products_references.WithProductsReferenceEntityRecordsListLabels("{}"),
    products_references.WithProductsReferenceEntityRecordsListAttributeValues("{}"),
    products_references.WithProductsReferenceEntityRecordsListCreatedAt("2026-01-01T12:00:00Z"),
    products_references.WithProductsReferenceEntityRecordsListUpdatedAt("2026-01-01T12:00:00Z"),
)
```
