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

response, error := service.CustomersLifecycleStagesUpdate(
    "",
    customers_value_lists.WithCustomersLifecycleStagesUpdateDescription("Has ordered at least once and is being served."),
    customers_value_lists.WithCustomersLifecycleStagesUpdateDescriptions(map[string]interface{}{
        "de": "Hat mindestens einmal bestellt und wird betreut.",
        "en": "Has ordered at least once and is being served."
    }),
    customers_value_lists.WithCustomersLifecycleStagesUpdateIsDefault(true),
    customers_value_lists.WithCustomersLifecycleStagesUpdateLabels(map[string]interface{}{
        "de": "Kunde",
        "en": "Customer"
    }),
    customers_value_lists.WithCustomersLifecycleStagesUpdatePosition(1),
    customers_value_lists.WithCustomersLifecycleStagesUpdateTitle("Customer"),
    customers_value_lists.WithCustomersLifecycleStagesUpdateTone("neutral"),
)
```
