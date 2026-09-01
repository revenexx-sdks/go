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

response, error := service.MarketsCreate(
    "northwind",
    "Northwind",
    markets.WithMarketsCreateCurrency("EUR"),
    markets.WithMarketsCreateIsDefault(false),
    markets.WithMarketsCreateLabels(map[string]interface{}{
        "de-DE": "Nordwind",
        "en-GB": "Northwind"
    }),
    markets.WithMarketsCreatePosition(0),
    markets.WithMarketsCreateStatus("active"),
)
```
