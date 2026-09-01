```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/payments_methods"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := payments_methods.New(client)

response, error := service.PaymentsMethodsEligible(
    payments_methods.WithPaymentsMethodsEligibleAmount(49.9),
    payments_methods.WithPaymentsMethodsEligibleCountry("DE"),
    payments_methods.WithPaymentsMethodsEligibleCurrency("EUR"),
)
```
