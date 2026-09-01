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

response, error := service.PagesDeliveryPage(
    pages_delivery.WithPagesDeliveryPageSlug("about-us"),
    pages_delivery.WithPagesDeliveryPageId(""),
    pages_delivery.WithPagesDeliveryPageLangcode("de"),
)
```
