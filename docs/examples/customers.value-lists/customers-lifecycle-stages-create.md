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

response, error := service.CustomersLifecycleStagesCreate(
    "",
    "Customer",
    customers_value_lists.WithCustomersLifecycleStagesCreateDescription("Has ordered at least once and is being served."),
    customers_value_lists.WithCustomersLifecycleStagesCreateDescriptions(map[string]interface{}{
        "de": "Hat mindestens einmal bestellt und wird betreut.",
        "en": "Has ordered at least once and is being served."
    }),
    customers_value_lists.WithCustomersLifecycleStagesCreateIsDefault(true),
    customers_value_lists.WithCustomersLifecycleStagesCreateLabels(map[string]interface{}{
        "de": "Kunde",
        "en": "Customer"
    }),
    customers_value_lists.WithCustomersLifecycleStagesCreatePosition(1),
    customers_value_lists.WithCustomersLifecycleStagesCreateTone("neutral"),
)
```
