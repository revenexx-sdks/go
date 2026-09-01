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

response, error := service.MarketsList(
    markets.WithMarketsListId(""),
    markets.WithMarketsListCode("northwind"),
    markets.WithMarketsListName("Northwind"),
    markets.WithMarketsListLabels("{"de-DE":"Nordwind","en-GB":"Northwind"}"),
    markets.WithMarketsListCurrency("EUR"),
    markets.WithMarketsListStatus("active"),
    markets.WithMarketsListIsDefault(false),
    markets.WithMarketsListPosition(0),
    markets.WithMarketsListCreatedAt("2026-01-01T12:00:00Z"),
    markets.WithMarketsListUpdatedAt("2026-01-01T12:00:00Z"),
    markets.WithMarketsListLimit(50),
    markets.WithMarketsListOffset(0),
    markets.WithMarketsListOrder("position.asc"),
)
```
