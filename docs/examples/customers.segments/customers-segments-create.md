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

response, error := service.CustomersSegmentsCreate(
    "key_accounts",
    customers_segments.WithCustomersSegmentsCreateLabels(map[string]interface{}{
        "de": "Gro\u00dfkunden",
        "en": "Key accounts"
    }),
    customers_segments.WithCustomersSegmentsCreatePosition(1),
    customers_segments.WithCustomersSegmentsCreateRuleMatch("all"),
    customers_segments.WithCustomersSegmentsCreateRules(map[string]interface{}{}),
)
```
