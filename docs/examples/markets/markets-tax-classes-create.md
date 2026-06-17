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

response, error := service.MarketsTaxClassesCreate(
    "",
    "",
    "",
    markets.WithMarketsTaxClassesCreateIsDefault(false),
    markets.WithMarketsTaxClassesCreateLabels(map[string]interface{}{}),
    markets.WithMarketsTaxClassesCreatePosition(0),
    markets.WithMarketsTaxClassesCreateRate(0),
)
```
