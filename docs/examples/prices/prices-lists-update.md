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

response, error := service.PricesListsUpdate(
    "",
    prices.WithPricesListsUpdateChannelId(""),
    prices.WithPricesListsUpdateCode("dealer-de"),
    prices.WithPricesListsUpdateContactId(""),
    prices.WithPricesListsUpdateCurrency("EUR"),
    prices.WithPricesListsUpdateDescription("Contract prices for authorised dealers."),
    prices.WithPricesListsUpdateIsDefault(true),
    prices.WithPricesListsUpdateLabels(map[string]interface{}{
        "de": "H\u00e4ndlerpreise",
        "en": "Dealer prices"
    }),
    prices.WithPricesListsUpdateMetadata(map[string]interface{}{
        "erp_price_group": "A1",
        "source_system": "erp"
    }),
    prices.WithPricesListsUpdateName("Dealer prices"),
    prices.WithPricesListsUpdateOrganizationId(""),
    prices.WithPricesListsUpdatePriority(1),
    prices.WithPricesListsUpdateRequiresAuth(true),
    prices.WithPricesListsUpdateStatus("active"),
    prices.WithPricesListsUpdateTaxBasis("net"),
    prices.WithPricesListsUpdateTaxIncluded(true),
    prices.WithPricesListsUpdateValidFrom("2026-01-01T00:00:00Z"),
    prices.WithPricesListsUpdateValidUntil("2026-12-31T23:59:59Z"),
)
```
