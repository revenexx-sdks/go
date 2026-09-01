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

response, error := service.MarketsClone(
    "northwind",
    "northwind-b2b",
    markets.WithMarketsCloneCopyCurrencies(true),
    markets.WithMarketsCloneCopyLocales(true),
    markets.WithMarketsCloneCopyTaxClasses(true),
    markets.WithMarketsCloneCurrency("EUR"),
    markets.WithMarketsCloneName("Northwind B2B"),
    markets.WithMarketsCloneStatus("active"),
)
```
