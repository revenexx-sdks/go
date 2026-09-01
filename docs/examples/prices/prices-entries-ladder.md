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

response, error := service.PricesEntriesLadder(
    "",
    9.99,
    prices.WithPricesEntriesLadderDiscountPercent(9.99),
    prices.WithPricesEntriesLadderProductId(""),
    prices.WithPricesEntriesLadderQuantities(interface{}{1,10,50}),
    prices.WithPricesEntriesLadderReplace(true),
    prices.WithPricesEntriesLadderRounding("exact"),
    prices.WithPricesEntriesLadderSku("BOLT-M8-30"),
    prices.WithPricesEntriesLadderUnit("pcs"),
)
```
