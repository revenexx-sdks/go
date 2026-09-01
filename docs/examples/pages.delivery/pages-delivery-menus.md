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

response, error := service.PagesDeliveryMenus(
    pages_delivery.WithPagesDeliveryMenusLimit(1),
    pages_delivery.WithPagesDeliveryMenusOffset(1),
    pages_delivery.WithPagesDeliveryMenusOrder("created_at.desc"),
)
```
