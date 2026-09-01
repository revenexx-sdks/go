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

response, error := service.MarketsUpdate(
    "",
    markets.WithMarketsUpdateCode("northwind"),
    markets.WithMarketsUpdateCurrency("EUR"),
    markets.WithMarketsUpdateIsDefault(false),
    markets.WithMarketsUpdateLabels(map[string]interface{}{
        "de-DE": "Nordwind",
        "en-GB": "Northwind"
    }),
    markets.WithMarketsUpdateName("Northwind"),
    markets.WithMarketsUpdatePosition(0),
    markets.WithMarketsUpdateStatus("active"),
)
```
