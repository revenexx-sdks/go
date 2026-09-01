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

response, error := service.PricesEntriesUpdate(
    "",
    "",
    prices.WithPricesEntriesUpdateMetadata(map[string]interface{}{
        "imported_batch": "2026-02-14",
        "source_system": "erp"
    }),
    prices.WithPricesEntriesUpdatePriceType("standard"),
    prices.WithPricesEntriesUpdateProductId(""),
    prices.WithPricesEntriesUpdateQuantityMin(9.99),
    prices.WithPricesEntriesUpdateSku("BOLT-M8-30"),
    prices.WithPricesEntriesUpdateUnit("pcs"),
    prices.WithPricesEntriesUpdateUnitPrice(9.99),
    prices.WithPricesEntriesUpdateValidFrom("2026-03-01T00:00:00Z"),
    prices.WithPricesEntriesUpdateValidUntil("2026-03-31T23:59:59Z"),
)
```
