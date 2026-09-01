```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/avatars"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := avatars.New(client)

response, error := service.AvatarsGetScreenshot(
    "https://example.com",
    avatars.WithAvatarsGetScreenshotHeaders(map[string]interface{}{}),
    avatars.WithAvatarsGetScreenshotViewportWidth(1),
    avatars.WithAvatarsGetScreenshotViewportHeight(1),
    avatars.WithAvatarsGetScreenshotScale(1),
    avatars.WithAvatarsGetScreenshotTheme("light"),
    avatars.WithAvatarsGetScreenshotUserAgent("Mozilla/5.0 (iPhone; CPU iPhone OS 14_0 like Mac OS X) AppleWebKit/605.1.15"),
    avatars.WithAvatarsGetScreenshotFullpage(true),
    avatars.WithAvatarsGetScreenshotLocale("en-US"),
    avatars.WithAvatarsGetScreenshotTimezone("Africa/Abidjan"),
    avatars.WithAvatarsGetScreenshotLatitude(9.99),
    avatars.WithAvatarsGetScreenshotLongitude(9.99),
    avatars.WithAvatarsGetScreenshotAccuracy(9.99),
    avatars.WithAvatarsGetScreenshotTouch(true),
    avatars.WithAvatarsGetScreenshotPermissions([]interface{}{}),
    avatars.WithAvatarsGetScreenshotSleep(1),
    avatars.WithAvatarsGetScreenshotWidth(1),
    avatars.WithAvatarsGetScreenshotHeight(1),
    avatars.WithAvatarsGetScreenshotQuality(1),
    avatars.WithAvatarsGetScreenshotOutput("jpg"),
)
```
