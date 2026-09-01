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

response, error := service.MarketsCurrenciesList(
    "",
    markets.WithMarketsCurrenciesListId(""),
    markets.WithMarketsCurrenciesListCode("EUR"),
    markets.WithMarketsCurrenciesListIsDefault(true),
    markets.WithMarketsCurrenciesListPosition(0),
    markets.WithMarketsCurrenciesListCreatedAt("2026-01-01T12:00:00Z"),
    markets.WithMarketsCurrenciesListLimit(50),
    markets.WithMarketsCurrenciesListOffset(0),
    markets.WithMarketsCurrenciesListOrder("position.asc"),
)
```
