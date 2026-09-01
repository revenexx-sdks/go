```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/markets"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := markets.New(client)

response, error := service.MarketsBackfill(
    "northwind",
    "northwind",
    markets.WithMarketsBackfillCurrencies(true),
    markets.WithMarketsBackfillLocales(true),
    markets.WithMarketsBackfillTaxClasses(true),
)
```
