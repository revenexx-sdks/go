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

response, error := service.PaymentsMethodsEligible(
    payments.WithPaymentsMethodsEligibleAmount(0),
    payments.WithPaymentsMethodsEligibleCountry(""),
    payments.WithPaymentsMethodsEligibleCurrency(""),
)
```
