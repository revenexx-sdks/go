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

response, error := service.CustomersSegmentMembersCreate(
    "",
    "",
    customers_segments.WithCustomersSegmentMembersCreateSource("manual"),
)
```
