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

response, error := service.PaymentsProvidersUpdate(
    "",
    payments_providers.WithPaymentsProvidersUpdateCredentials(map[string]interface{}{}),
    payments_providers.WithPaymentsProvidersUpdateEnabled(true),
    payments_providers.WithPaymentsProvidersUpdateName("Stripe"),
    payments_providers.WithPaymentsProvidersUpdateOptions(map[string]interface{}{
        "capture_method": "automatic",
        "logo_url": "https:\/\/apps.example.com\/payments\/logos\/stripe",
        "three_ds": false
    }),
    payments_providers.WithPaymentsProvidersUpdateProvider("stripe"),
    payments_providers.WithPaymentsProvidersUpdateTestMode(true),
    payments_providers.WithPaymentsProvidersUpdateWebhookSecret(""),
)
```
