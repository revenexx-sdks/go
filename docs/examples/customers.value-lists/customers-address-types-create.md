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

response, error := service.CustomersAddressTypesCreate(
    "",
    "Shipping address",
    customers_value_lists.WithCustomersAddressTypesCreateDescription("Where the goods go."),
    customers_value_lists.WithCustomersAddressTypesCreateDescriptions(map[string]interface{}{
        "de": "Wohin die Ware geliefert wird.",
        "en": "Where the goods go."
    }),
    customers_value_lists.WithCustomersAddressTypesCreateIsDefault(true),
    customers_value_lists.WithCustomersAddressTypesCreateLabels(map[string]interface{}{
        "de": "Lieferadresse",
        "en": "Shipping address"
    }),
    customers_value_lists.WithCustomersAddressTypesCreatePosition(1),
    customers_value_lists.WithCustomersAddressTypesCreateTone("neutral"),
)
```
