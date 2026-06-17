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

response, error := service.MarketsCreate(
    "",
    "",
    markets.WithMarketsCreateCurrency(""),
    markets.WithMarketsCreateIsDefault(false),
    markets.WithMarketsCreateLabels(map[string]interface{}{}),
    markets.WithMarketsCreatePosition(0),
    markets.WithMarketsCreateStatus(""),
)
```
