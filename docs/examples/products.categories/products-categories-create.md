```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/products_categories"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := products_categories.New(client)

response, error := service.ProductsCategoriesCreate(
    "cordless_drills",
    products_categories.WithProductsCategoriesCreateLabels(map[string]interface{}{
        "de": "Akku-Bohrschrauber",
        "en": "Cordless drills"
    }),
    products_categories.WithProductsCategoriesCreateParentId(""),
    products_categories.WithProductsCategoriesCreatePath("tools/power_tools/cordless_drills"),
    products_categories.WithProductsCategoriesCreatePosition(1),
    products_categories.WithProductsCategoriesCreateRuleMatch("all"),
    products_categories.WithProductsCategoriesCreateRules(map[string]interface{}{
        "conditions": [
            {
                "field": "attribute:brand",
                "operator": "in",
                "value": [
                    "acme",
                    "globex"
                ]
            },
            {
                "field": "enabled",
                "operator": "eq",
                "value": true
            }
        ]
    }),
    products_categories.WithProductsCategoriesCreateRulesComputedAt("2026-01-01T12:00:00Z"),
    products_categories.WithProductsCategoriesCreateValues(map[string]interface{}{
        "hero_asset": "packshots\/cordless_drills_hero",
        "seo_title": "Cordless drills"
    }),
)
```
