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

response, error := service.PricesEntriesCreate(
    "",
    prices.WithPricesEntriesCreateMetadata(map[string]interface{}{
        "imported_batch": "2026-02-14",
        "source_system": "erp"
    }),
    prices.WithPricesEntriesCreatePriceType("standard"),
    prices.WithPricesEntriesCreateProductId(""),
    prices.WithPricesEntriesCreateQuantityMin(9.99),
    prices.WithPricesEntriesCreateSku("BOLT-M8-30"),
    prices.WithPricesEntriesCreateUnit("pcs"),
    prices.WithPricesEntriesCreateUnitPrice(9.99),
    prices.WithPricesEntriesCreateValidFrom("2026-03-01T00:00:00Z"),
    prices.WithPricesEntriesCreateValidUntil("2026-03-31T23:59:59Z"),
)
```
