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

response, error := service.ProductsCategoriesList(
    products_categories.WithProductsCategoriesListLimit(1),
    products_categories.WithProductsCategoriesListOffset(1),
    products_categories.WithProductsCategoriesListOrder("created_at.desc"),
    products_categories.WithProductsCategoriesListId(""),
    products_categories.WithProductsCategoriesListCode("cordless_drills"),
    products_categories.WithProductsCategoriesListParentId(""),
    products_categories.WithProductsCategoriesListPath("tools/power_tools/cordless_drills"),
    products_categories.WithProductsCategoriesListPosition(1),
    products_categories.WithProductsCategoriesListLabels("{}"),
    products_categories.WithProductsCategoriesListValues("{}"),
    products_categories.WithProductsCategoriesListRules("{}"),
    products_categories.WithProductsCategoriesListRuleMatch("all"),
    products_categories.WithProductsCategoriesListRulesComputedAt("2026-01-01T12:00:00Z"),
    products_categories.WithProductsCategoriesListCreatedAt("2026-01-01T12:00:00Z"),
    products_categories.WithProductsCategoriesListUpdatedAt("2026-01-01T12:00:00Z"),
)
```
