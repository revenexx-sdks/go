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

response, error := service.PricesEntriesCreate(
    "",
    prices.WithPricesEntriesCreateMetadata(map[string]interface{}{}),
    prices.WithPricesEntriesCreatePriceType(""),
    prices.WithPricesEntriesCreateProductId(""),
    prices.WithPricesEntriesCreateQuantityMin(0),
    prices.WithPricesEntriesCreateSku(""),
    prices.WithPricesEntriesCreateUnit(""),
    prices.WithPricesEntriesCreateUnitPrice(0),
    prices.WithPricesEntriesCreateValidFrom(""),
    prices.WithPricesEntriesCreateValidUntil(""),
)
```
