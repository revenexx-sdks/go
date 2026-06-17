```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/markets"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com")
    client.WithApiKeyAuth("<API_KEY>")
)

service := markets.New(client)

response, error := service.MarketsTaxClassesUpdate(
    "",
    "",
    markets.WithMarketsTaxClassesUpdateCode(""),
    markets.WithMarketsTaxClassesUpdateIsDefault(false),
    markets.WithMarketsTaxClassesUpdateLabels(map[string]interface{}{}),
    markets.WithMarketsTaxClassesUpdateName(""),
    markets.WithMarketsTaxClassesUpdatePosition(0),
    markets.WithMarketsTaxClassesUpdateRate(0),
)
```
