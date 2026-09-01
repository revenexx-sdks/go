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

response, error := service.MarketsTaxClassesList(
    "",
    markets.WithMarketsTaxClassesListId(""),
    markets.WithMarketsTaxClassesListCode("standard"),
    markets.WithMarketsTaxClassesListName("Standard rate"),
    markets.WithMarketsTaxClassesListLabels("{"de-DE":"Regelsatz","en-GB":"Standard rate"}"),
    markets.WithMarketsTaxClassesListRate(20),
    markets.WithMarketsTaxClassesListIsDefault(true),
    markets.WithMarketsTaxClassesListPosition(0),
    markets.WithMarketsTaxClassesListCreatedAt("2026-01-01T12:00:00Z"),
    markets.WithMarketsTaxClassesListUpdatedAt("2026-01-01T12:00:00Z"),
    markets.WithMarketsTaxClassesListLimit(50),
    markets.WithMarketsTaxClassesListOffset(0),
    markets.WithMarketsTaxClassesListOrder("position.asc"),
)
```
