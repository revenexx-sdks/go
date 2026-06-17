```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/payments"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com")
    client.WithApiKeyAuth("<API_KEY>")
)

service := payments.New(client)

response, error := service.PaymentsProvidersCreate(
    "",
    payments.WithPaymentsProvidersCreateCredentials(map[string]interface{}{}),
    payments.WithPaymentsProvidersCreateEnabled(false),
    payments.WithPaymentsProvidersCreateName(""),
    payments.WithPaymentsProvidersCreateOptions(map[string]interface{}{}),
    payments.WithPaymentsProvidersCreateTestMode(false),
    payments.WithPaymentsProvidersCreateWebhookSecret(""),
)
```
