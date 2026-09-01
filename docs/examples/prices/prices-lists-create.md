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

response, error := service.PricesListsCreate(
    "dealer-de",
    "Dealer prices",
    prices.WithPricesListsCreateChannelId(""),
    prices.WithPricesListsCreateContactId(""),
    prices.WithPricesListsCreateCurrency("EUR"),
    prices.WithPricesListsCreateDescription("Contract prices for authorised dealers."),
    prices.WithPricesListsCreateIsDefault(true),
    prices.WithPricesListsCreateLabels(map[string]interface{}{
        "de": "H\u00e4ndlerpreise",
        "en": "Dealer prices"
    }),
    prices.WithPricesListsCreateMetadata(map[string]interface{}{
        "erp_price_group": "A1",
        "source_system": "erp"
    }),
    prices.WithPricesListsCreateOrganizationId(""),
    prices.WithPricesListsCreatePriority(1),
    prices.WithPricesListsCreateRequiresAuth(true),
    prices.WithPricesListsCreateStatus("active"),
    prices.WithPricesListsCreateTaxBasis("net"),
    prices.WithPricesListsCreateTaxIncluded(true),
    prices.WithPricesListsCreateValidFrom("2026-01-01T00:00:00Z"),
    prices.WithPricesListsCreateValidUntil("2026-12-31T23:59:59Z"),
)
```
