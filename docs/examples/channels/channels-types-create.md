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

response, error := service.ChannelsTypesCreate(
    "feed",
    "Product feed",
    channels.WithChannelsTypesCreateDescription("A web shop a human browses."),
    channels.WithChannelsTypesCreateDescriptions(map[string]interface{}{
        "de": "Shop",
        "en": "Shop"
    }),
    channels.WithChannelsTypesCreateIsDefault(true),
    channels.WithChannelsTypesCreateLabels(map[string]interface{}{
        "de": "Shop",
        "en": "Shop"
    }),
    channels.WithChannelsTypesCreatePosition(1),
    channels.WithChannelsTypesCreateTone("neutral"),
)
```
