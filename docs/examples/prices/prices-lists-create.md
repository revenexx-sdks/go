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

response, error := service.PricesListsCreate(
    "",
    "",
    prices.WithPricesListsCreateChannelId(""),
    prices.WithPricesListsCreateContactId(""),
    prices.WithPricesListsCreateCurrency(""),
    prices.WithPricesListsCreateDescription(""),
    prices.WithPricesListsCreateIsDefault(false),
    prices.WithPricesListsCreateLabels(map[string]interface{}{}),
    prices.WithPricesListsCreateMarketId(""),
    prices.WithPricesListsCreateMetadata(map[string]interface{}{}),
    prices.WithPricesListsCreateOrganizationId(""),
    prices.WithPricesListsCreatePriority(0),
    prices.WithPricesListsCreateStatus(""),
    prices.WithPricesListsCreateTaxIncluded(false),
    prices.WithPricesListsCreateValidFrom(""),
    prices.WithPricesListsCreateValidUntil(""),
)
```
