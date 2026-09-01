```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/prices"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := prices.New(client)

response, error := service.PricesResolve(
    []interface{}{},
    prices.WithPricesResolveAt("2026-03-15T09:00:00Z"),
    prices.WithPricesResolveChannelId(""),
    prices.WithPricesResolveContactId(""),
    prices.WithPricesResolveCurrency("EUR"),
    prices.WithPricesResolveMarketId(""),
    prices.WithPricesResolveOrganizationId(""),
)
```
