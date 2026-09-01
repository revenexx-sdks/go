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

response, error := service.ChannelsUpdate(
    "",
    channels.WithChannelsUpdateCode("shop"),
    channels.WithChannelsUpdateIsDefault(true),
    channels.WithChannelsUpdateLabels(map[string]interface{}{
        "de": "Shop",
        "en": "Shop"
    }),
    channels.WithChannelsUpdateName("Shop"),
    channels.WithChannelsUpdatePosition(1),
    channels.WithChannelsUpdateStatus("active"),
    channels.WithChannelsUpdateType("storefront"),
    channels.WithChannelsUpdateUnassignedVisibility("inherit"),
)
```
