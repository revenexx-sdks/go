```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/prices"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com")
    client.WithApiKeyAuth("<API_KEY>")
)

service := prices.New(client)

response, error := service.PricesResolve(
    []interface{}{},
    prices.WithPricesResolveAt(""),
    prices.WithPricesResolveChannelId(""),
    prices.WithPricesResolveContactId(""),
    prices.WithPricesResolveCurrency(""),
    prices.WithPricesResolveMarketId(""),
    prices.WithPricesResolveOrganizationId(""),
)
```
