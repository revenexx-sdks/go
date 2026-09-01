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

response, error := service.ProductsCategoriesUpdate(
    "",
    products_categories.WithProductsCategoriesUpdateCode("cordless_drills"),
    products_categories.WithProductsCategoriesUpdateLabels(map[string]interface{}{
        "de": "Akku-Bohrschrauber",
        "en": "Cordless drills"
    }),
    products_categories.WithProductsCategoriesUpdateParentId(""),
    products_categories.WithProductsCategoriesUpdatePath("tools/power_tools/cordless_drills"),
    products_categories.WithProductsCategoriesUpdatePosition(1),
    products_categories.WithProductsCategoriesUpdateRuleMatch("all"),
    products_categories.WithProductsCategoriesUpdateRules(map[string]interface{}{
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
    products_categories.WithProductsCategoriesUpdateRulesComputedAt("2026-01-01T12:00:00Z"),
    products_categories.WithProductsCategoriesUpdateValues(map[string]interface{}{
        "hero_asset": "packshots\/cordless_drills_hero",
        "seo_title": "Cordless drills"
    }),
)
```
