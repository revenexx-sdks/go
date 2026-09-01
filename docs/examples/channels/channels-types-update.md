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

response, error := service.ChannelsTypesUpdate(
    "",
    channels.WithChannelsTypesUpdateDescription("A web shop a human browses."),
    channels.WithChannelsTypesUpdateDescriptions(map[string]interface{}{
        "de": "Shop",
        "en": "Shop"
    }),
    channels.WithChannelsTypesUpdateIsDefault(true),
    channels.WithChannelsTypesUpdateLabels(map[string]interface{}{
        "de": "Shop",
        "en": "Shop"
    }),
    channels.WithChannelsTypesUpdatePosition(1),
    channels.WithChannelsTypesUpdateTitle("Product feed"),
    channels.WithChannelsTypesUpdateTone("neutral"),
)
```
