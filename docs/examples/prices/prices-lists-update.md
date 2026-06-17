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

response, error := service.PricesListsUpdate(
    "",
    prices.WithPricesListsUpdateChannelId(""),
    prices.WithPricesListsUpdateCode(""),
    prices.WithPricesListsUpdateContactId(""),
    prices.WithPricesListsUpdateCurrency(""),
    prices.WithPricesListsUpdateDescription(""),
    prices.WithPricesListsUpdateIsDefault(false),
    prices.WithPricesListsUpdateLabels(map[string]interface{}{}),
    prices.WithPricesListsUpdateMarketId(""),
    prices.WithPricesListsUpdateMetadata(map[string]interface{}{}),
    prices.WithPricesListsUpdateName(""),
    prices.WithPricesListsUpdateOrganizationId(""),
    prices.WithPricesListsUpdatePriority(0),
    prices.WithPricesListsUpdateStatus(""),
    prices.WithPricesListsUpdateTaxIncluded(false),
    prices.WithPricesListsUpdateValidFrom(""),
    prices.WithPricesListsUpdateValidUntil(""),
)
```
