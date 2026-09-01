```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/orderlists"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := orderlists.New(client)

response, error := service.OrderlistsKindsUpdate(
    "",
    orderlists.WithOrderlistsKindsUpdateDescription("Chemicals ordered against a standing lab protocol."),
    orderlists.WithOrderlistsKindsUpdateDescriptions(map[string]interface{}{
        "de": "Chemikalien, die nach einem festen Laborprotokoll bestellt werden.",
        "en": "Chemicals ordered against a standing lab protocol."
    }),
    orderlists.WithOrderlistsKindsUpdateIsDefault(true),
    orderlists.WithOrderlistsKindsUpdateLabels(map[string]interface{}{
        "de": "Reagenzienliste",
        "en": "Reagent list"
    }),
    orderlists.WithOrderlistsKindsUpdatePosition(2),
    orderlists.WithOrderlistsKindsUpdateTitle("Reagent list"),
    orderlists.WithOrderlistsKindsUpdateTone("neutral"),
)
```
