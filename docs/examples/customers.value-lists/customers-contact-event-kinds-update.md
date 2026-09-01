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

response, error := service.CustomersContactEventKindsUpdate(
    "",
    customers_value_lists.WithCustomersContactEventKindsUpdateDescription("Somebody spoke to this person on the phone."),
    customers_value_lists.WithCustomersContactEventKindsUpdateDescriptions(map[string]interface{}{
        "de": "Es wurde mit dieser Person telefoniert.",
        "en": "Somebody spoke to this person on the phone."
    }),
    customers_value_lists.WithCustomersContactEventKindsUpdateIsDefault(true),
    customers_value_lists.WithCustomersContactEventKindsUpdateLabels(map[string]interface{}{
        "de": "Telefonat",
        "en": "Phone call"
    }),
    customers_value_lists.WithCustomersContactEventKindsUpdatePosition(1),
    customers_value_lists.WithCustomersContactEventKindsUpdateTitle("Phone call"),
    customers_value_lists.WithCustomersContactEventKindsUpdateTone("neutral"),
)
```
