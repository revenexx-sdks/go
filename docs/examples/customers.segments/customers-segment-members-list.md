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

response, error := service.CustomersSegmentMembersList(
    customers_segments.WithCustomersSegmentMembersListId(""),
    customers_segments.WithCustomersSegmentMembersListSegmentId(""),
    customers_segments.WithCustomersSegmentMembersListOrganizationId(""),
    customers_segments.WithCustomersSegmentMembersListSource("manual"),
    customers_segments.WithCustomersSegmentMembersListCreatedAt("2026-01-01T12:00:00Z"),
    customers_segments.WithCustomersSegmentMembersListLimit(1),
    customers_segments.WithCustomersSegmentMembersListOffset(1),
    customers_segments.WithCustomersSegmentMembersListOrder("created_at.desc"),
)
```
