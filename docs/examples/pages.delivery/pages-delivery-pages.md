```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/pages_delivery"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := pages_delivery.New(client)

response, error := service.PagesDeliveryPages(
    pages_delivery.WithPagesDeliveryPagesLimit(1),
    pages_delivery.WithPagesDeliveryPagesOffset(1),
    pages_delivery.WithPagesDeliveryPagesOrder("created_at.desc"),
    pages_delivery.WithPagesDeliveryPagesBundle("standard"),
)
```
