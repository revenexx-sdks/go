```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/channels"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com")
    client.WithApiKeyAuth("<API_KEY>")
)

service := channels.New(client)

response, error := service.ChannelsUpdate(
    "",
    channels.WithChannelsUpdateCode(""),
    channels.WithChannelsUpdateIsDefault(false),
    channels.WithChannelsUpdateLabels(map[string]interface{}{}),
    channels.WithChannelsUpdateName(""),
    channels.WithChannelsUpdatePosition(0),
    channels.WithChannelsUpdateStatus(""),
    channels.WithChannelsUpdateType(""),
)
```
