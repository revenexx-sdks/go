```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/customers_value_lists"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := customers_value_lists.New(client)

response, error := service.CustomersAddressTypesUpdate(
    "",
    customers_value_lists.WithCustomersAddressTypesUpdateDescription("Where the goods go."),
    customers_value_lists.WithCustomersAddressTypesUpdateDescriptions(map[string]interface{}{
        "de": "Wohin die Ware geliefert wird.",
        "en": "Where the goods go."
    }),
    customers_value_lists.WithCustomersAddressTypesUpdateIsDefault(true),
    customers_value_lists.WithCustomersAddressTypesUpdateLabels(map[string]interface{}{
        "de": "Lieferadresse",
        "en": "Shipping address"
    }),
    customers_value_lists.WithCustomersAddressTypesUpdatePosition(1),
    customers_value_lists.WithCustomersAddressTypesUpdateTitle("Shipping address"),
    customers_value_lists.WithCustomersAddressTypesUpdateTone("neutral"),
)
```
