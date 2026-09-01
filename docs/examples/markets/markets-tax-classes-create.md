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

response, error := service.MarketsTaxClassesCreate(
    "",
    "standard",
    "Standard rate",
    markets.WithMarketsTaxClassesCreateIsDefault(true),
    markets.WithMarketsTaxClassesCreateLabels(map[string]interface{}{
        "de-DE": "Regelsatz",
        "en-GB": "Standard rate"
    }),
    markets.WithMarketsTaxClassesCreatePosition(0),
    markets.WithMarketsTaxClassesCreateRate(20),
)
```
