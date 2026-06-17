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

response, error := service.PaymentsCreate(
    0,
    "",
    payments.WithPaymentsCreateCartId(""),
    payments.WithPaymentsCreateContactId(""),
    payments.WithPaymentsCreateCountry(""),
    payments.WithPaymentsCreateCurrency(""),
    payments.WithPaymentsCreateIdempotencyKey(""),
    payments.WithPaymentsCreateMetadata(map[string]interface{}{}),
    payments.WithPaymentsCreateOrderRef(""),
    payments.WithPaymentsCreateReturnUrl(""),
)
```
