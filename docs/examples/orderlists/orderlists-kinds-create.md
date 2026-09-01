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

response, error := service.OrderlistsKindsCreate(
    "reagents",
    "Reagent list",
    orderlists.WithOrderlistsKindsCreateDescription("Chemicals ordered against a standing lab protocol."),
    orderlists.WithOrderlistsKindsCreateDescriptions(map[string]interface{}{
        "de": "Chemikalien, die nach einem festen Laborprotokoll bestellt werden.",
        "en": "Chemicals ordered against a standing lab protocol."
    }),
    orderlists.WithOrderlistsKindsCreateIsDefault(true),
    orderlists.WithOrderlistsKindsCreateLabels(map[string]interface{}{
        "de": "Reagenzienliste",
        "en": "Reagent list"
    }),
    orderlists.WithOrderlistsKindsCreatePosition(2),
    orderlists.WithOrderlistsKindsCreateTone("neutral"),
)
```
