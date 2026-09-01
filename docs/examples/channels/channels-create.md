```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/channels"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := channels.New(client)

response, error := service.ChannelsCreate(
    "shop",
    "Shop",
    channels.WithChannelsCreateIsDefault(true),
    channels.WithChannelsCreateLabels(map[string]interface{}{
        "de": "Shop",
        "en": "Shop"
    }),
    channels.WithChannelsCreatePosition(1),
    channels.WithChannelsCreateStatus("active"),
    channels.WithChannelsCreateType("storefront"),
    channels.WithChannelsCreateUnassignedVisibility("inherit"),
)
```
