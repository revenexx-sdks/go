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

response, error := service.PagesDeliveryPreview(
    "",
    pages_delivery.WithPagesDeliveryPreviewLangcode("de"),
)
```
