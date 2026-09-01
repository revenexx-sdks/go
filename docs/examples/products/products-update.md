```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/products"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := products.New(client)

response, error := service.ProductsUpdate(
    "",
    products.WithProductsUpdateAttributeValues(map[string]interface{}{
        "channel_locale_specific": {
            "b2b": {
                "de_DE": {
                    "description": "Staffelpreise auf Anfrage."
                }
            }
        },
        "channel_specific": {
            "b2b": {
                "minimum_order_quantity": 6
            }
        },
        "common": {
            "colour": "black",
            "manufacturer_aid": "4711-BLK",
            "net_weight": 2.4
        },
        "locale_specific": {
            "de_DE": {
                "description": "B\u00fcrstenloser Motor, 2 Akkus im Set.",
                "name": "Akku-Bohrschrauber 18V"
            },
            "en_GB": {
                "name": "18V cordless drill"
            }
        }
    }),
    products.WithProductsUpdateCompleteness(map[string]interface{}{
        "computed_at": "2026-01-01T12:00:00Z",
        "filled": 9,
        "missing": [
            "net_weight",
            "packaging_unit",
            "safety_datasheet"
        ],
        "ratio": 0.75,
        "required": 12
    }),
    products.WithProductsUpdateDeletedAt("2026-01-01T12:00:00Z"),
    products.WithProductsUpdateEnabled(true),
    products.WithProductsUpdateFamilyId(""),
    products.WithProductsUpdateFamilyVariantId(""),
    products.WithProductsUpdateKind("simple"),
    products.WithProductsUpdateParentId(""),
    products.WithProductsUpdateQuantifiedAssociations(map[string]interface{}{
        "PRODUCT_SET": {
            "product_models": [],
            "products": [
                {
                    "identifier": "ACME-4711-CASTER",
                    "quantity": 4
                }
            ]
        }
    }),
    products.WithProductsUpdateSku("ACME-4711-BLK"),
    products.WithProductsUpdateTaxClass("standard"),
)
```
