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

response, error := service.ChannelsList(
    channels.WithChannelsListId(""),
    channels.WithChannelsListCode("shop"),
    channels.WithChannelsListName("Shop"),
    channels.WithChannelsListLabels("{"en":"Shop","de":"Shop"}"),
    channels.WithChannelsListType("storefront"),
    channels.WithChannelsListStatus("active"),
    channels.WithChannelsListUnassignedVisibility("inherit"),
    channels.WithChannelsListIsDefault(true),
    channels.WithChannelsListPosition(1),
    channels.WithChannelsListCreatedAt("2026-01-01T12:00:00Z"),
    channels.WithChannelsListUpdatedAt("2026-01-01T12:00:00Z"),
    channels.WithChannelsListLimit(1),
    channels.WithChannelsListOffset(1),
    channels.WithChannelsListOrder("created_at.desc"),
)
```
