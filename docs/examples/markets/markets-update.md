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

response, error := service.MarketsUpdate(
    "",
    markets.WithMarketsUpdateCode(""),
    markets.WithMarketsUpdateCurrency(""),
    markets.WithMarketsUpdateIsDefault(false),
    markets.WithMarketsUpdateLabels(map[string]interface{}{}),
    markets.WithMarketsUpdateName(""),
    markets.WithMarketsUpdatePosition(0),
    markets.WithMarketsUpdateStatus(""),
)
```
