```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/prices"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := prices.New(client)

response, error := service.PricesEntriesList(
    "",
    prices.WithPricesEntriesListId(""),
    prices.WithPricesEntriesListProductId(""),
    prices.WithPricesEntriesListSku("BOLT-M8-30"),
    prices.WithPricesEntriesListPriceType("standard"),
    prices.WithPricesEntriesListQuantityMin(9.99),
    prices.WithPricesEntriesListUnitPrice(9.99),
    prices.WithPricesEntriesListUnit("pcs"),
    prices.WithPricesEntriesListValidFrom("2026-01-01T12:00:00Z"),
    prices.WithPricesEntriesListValidUntil("2026-01-01T12:00:00Z"),
    prices.WithPricesEntriesListCreatedAt("2026-01-01T12:00:00Z"),
    prices.WithPricesEntriesListUpdatedAt("2026-01-01T12:00:00Z"),
    prices.WithPricesEntriesListLimit(1),
    prices.WithPricesEntriesListOffset(1),
    prices.WithPricesEntriesListOrder("created_at.desc"),
)
```
