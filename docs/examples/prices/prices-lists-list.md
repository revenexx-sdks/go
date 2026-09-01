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

response, error := service.PricesListsList(
    prices.WithPricesListsListId(""),
    prices.WithPricesListsListCode("standard"),
    prices.WithPricesListsListName("Standard prices"),
    prices.WithPricesListsListDescription("The list every buyer falls back to."),
    prices.WithPricesListsListCurrency("EUR"),
    prices.WithPricesListsListStatus("active"),
    prices.WithPricesListsListPriority(1),
    prices.WithPricesListsListIsDefault(true),
    prices.WithPricesListsListTaxBasis("net"),
    prices.WithPricesListsListTaxIncluded(true),
    prices.WithPricesListsListRequiresAuth(true),
    prices.WithPricesListsListContactId(""),
    prices.WithPricesListsListOrganizationId(""),
    prices.WithPricesListsListChannelId(""),
    prices.WithPricesListsListValidFrom("2026-01-01T12:00:00Z"),
    prices.WithPricesListsListValidUntil("2026-01-01T12:00:00Z"),
    prices.WithPricesListsListCreatedAt("2026-01-01T12:00:00Z"),
    prices.WithPricesListsListUpdatedAt("2026-01-01T12:00:00Z"),
    prices.WithPricesListsListLimit(1),
    prices.WithPricesListsListOffset(1),
    prices.WithPricesListsListOrder("created_at.desc"),
)
```
