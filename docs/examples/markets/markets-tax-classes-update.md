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

response, error := service.MarketsTaxClassesUpdate(
    "",
    "",
    markets.WithMarketsTaxClassesUpdateCode("standard"),
    markets.WithMarketsTaxClassesUpdateIsDefault(true),
    markets.WithMarketsTaxClassesUpdateLabels(map[string]interface{}{
        "de-DE": "Regelsatz",
        "en-GB": "Standard rate"
    }),
    markets.WithMarketsTaxClassesUpdateName("Standard rate"),
    markets.WithMarketsTaxClassesUpdatePosition(0),
    markets.WithMarketsTaxClassesUpdateRate(20),
)
```
