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

response, error := service.CustomersSegmentsList(
    customers_segments.WithCustomersSegmentsListId(""),
    customers_segments.WithCustomersSegmentsListCode("key_accounts"),
    customers_segments.WithCustomersSegmentsListPosition(1),
    customers_segments.WithCustomersSegmentsListRuleMatch("all"),
    customers_segments.WithCustomersSegmentsListRulesComputedAt("2026-01-01T12:00:00Z"),
    customers_segments.WithCustomersSegmentsListCreatedAt("2026-01-01T12:00:00Z"),
    customers_segments.WithCustomersSegmentsListUpdatedAt("2026-01-01T12:00:00Z"),
    customers_segments.WithCustomersSegmentsListLimit(1),
    customers_segments.WithCustomersSegmentsListOffset(1),
    customers_segments.WithCustomersSegmentsListOrder("created_at.desc"),
)
```
