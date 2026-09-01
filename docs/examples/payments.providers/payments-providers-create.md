```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/payments_providers"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := payments_providers.New(client)

response, error := service.PaymentsProvidersCreate(
    "stripe",
    payments_providers.WithPaymentsProvidersCreateCredentials(map[string]interface{}{}),
    payments_providers.WithPaymentsProvidersCreateEnabled(true),
    payments_providers.WithPaymentsProvidersCreateName("Stripe"),
    payments_providers.WithPaymentsProvidersCreateOptions(map[string]interface{}{
        "capture_method": "automatic",
        "logo_url": "https:\/\/apps.example.com\/payments\/logos\/stripe",
        "three_ds": false
    }),
    payments_providers.WithPaymentsProvidersCreateTestMode(true),
    payments_providers.WithPaymentsProvidersCreateWebhookSecret(""),
)
```
