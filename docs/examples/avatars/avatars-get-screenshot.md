```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/avatars"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com")
    client.WithApiKeyAuth("<API_KEY>")
)

service := avatars.New(client)

response, error := service.AvatarsGetScreenshot(
    "",
    avatars.WithAvatarsGetScreenshotHeaders(map[string]interface{}{}),
    avatars.WithAvatarsGetScreenshotViewportWidth(0),
    avatars.WithAvatarsGetScreenshotViewportHeight(0),
    avatars.WithAvatarsGetScreenshotScale(0),
    avatars.WithAvatarsGetScreenshotTheme(""),
    avatars.WithAvatarsGetScreenshotUserAgent(""),
    avatars.WithAvatarsGetScreenshotFullpage(false),
    avatars.WithAvatarsGetScreenshotLocale(""),
    avatars.WithAvatarsGetScreenshotTimezone(""),
    avatars.WithAvatarsGetScreenshotLatitude(0),
    avatars.WithAvatarsGetScreenshotLongitude(0),
    avatars.WithAvatarsGetScreenshotAccuracy(0),
    avatars.WithAvatarsGetScreenshotTouch(false),
    avatars.WithAvatarsGetScreenshotPermissions([]interface{}{}),
    avatars.WithAvatarsGetScreenshotSleep(0),
    avatars.WithAvatarsGetScreenshotWidth(0),
    avatars.WithAvatarsGetScreenshotHeight(0),
    avatars.WithAvatarsGetScreenshotQuality(0),
    avatars.WithAvatarsGetScreenshotOutput(""),
)
```
