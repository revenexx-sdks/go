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

response, error := service.ProductsReferenceEntityRecordsCreate(
    "acme_tools",
    "",
    products_references.WithProductsReferenceEntityRecordsCreateAttributeValues(map[string]interface{}{
        "common": {
            "country": "DE",
            "founded": 1946
        },
        "locale_specific": {
            "de_DE": {
                "description": "Werkzeughersteller aus S\u00fcddeutschland."
            }
        }
    }),
    products_references.WithProductsReferenceEntityRecordsCreateLabels(map[string]interface{}{
        "de": "Acme Tools",
        "en": "Acme Tools"
    }),
)
```
