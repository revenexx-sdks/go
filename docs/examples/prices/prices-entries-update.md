```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/prices"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com")
    client.WithApiKeyAuth("<API_KEY>")
)

service := prices.New(client)

response, error := service.PricesEntriesUpdate(
    "",
    "",
    prices.WithPricesEntriesUpdateMetadata(map[string]interface{}{}),
    prices.WithPricesEntriesUpdatePriceType(""),
    prices.WithPricesEntriesUpdateProductId(""),
    prices.WithPricesEntriesUpdateQuantityMin(0),
    prices.WithPricesEntriesUpdateSku(""),
    prices.WithPricesEntriesUpdateUnit(""),
    prices.WithPricesEntriesUpdateUnitPrice(0),
    prices.WithPricesEntriesUpdateValidFrom(""),
    prices.WithPricesEntriesUpdateValidUntil(""),
)
```
