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

response, error := service.MarketsLocalesList(
    "",
    markets.WithMarketsLocalesListId(""),
    markets.WithMarketsLocalesListCode("de-DE"),
    markets.WithMarketsLocalesListLanguage("de"),
    markets.WithMarketsLocalesListCountry("DE"),
    markets.WithMarketsLocalesListIsDefault(true),
    markets.WithMarketsLocalesListPosition(0),
    markets.WithMarketsLocalesListCreatedAt("2026-01-01T12:00:00Z"),
    markets.WithMarketsLocalesListLimit(50),
    markets.WithMarketsLocalesListOffset(0),
    markets.WithMarketsLocalesListOrder("position.asc"),
)
```
