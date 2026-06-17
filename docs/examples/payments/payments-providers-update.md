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

response, error := service.PaymentsProvidersUpdate(
    "",
    payments.WithPaymentsProvidersUpdateCredentials(map[string]interface{}{}),
    payments.WithPaymentsProvidersUpdateEnabled(false),
    payments.WithPaymentsProvidersUpdateName(""),
    payments.WithPaymentsProvidersUpdateOptions(map[string]interface{}{}),
    payments.WithPaymentsProvidersUpdateProvider(""),
    payments.WithPaymentsProvidersUpdateTestMode(false),
    payments.WithPaymentsProvidersUpdateWebhookSecret(""),
)
```
