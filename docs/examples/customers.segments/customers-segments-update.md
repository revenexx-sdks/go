```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/customers_segments"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := customers_segments.New(client)

response, error := service.CustomersSegmentsUpdate(
    "",
    customers_segments.WithCustomersSegmentsUpdateCode("key_accounts"),
    customers_segments.WithCustomersSegmentsUpdateLabels(map[string]interface{}{
        "de": "Gro\u00dfkunden",
        "en": "Key accounts"
    }),
    customers_segments.WithCustomersSegmentsUpdatePosition(1),
    customers_segments.WithCustomersSegmentsUpdateRuleMatch("all"),
    customers_segments.WithCustomersSegmentsUpdateRules(map[string]interface{}{}),
)
```
