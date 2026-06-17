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

response, error := service.PaymentsMethodsCreate(
    "",
    "",
    payments.WithPaymentsMethodsCreateCountries([]interface{}{}),
    payments.WithPaymentsMethodsCreateDescription(""),
    payments.WithPaymentsMethodsCreateEnabled(false),
    payments.WithPaymentsMethodsCreateFeeAmount(0),
    payments.WithPaymentsMethodsCreateFeeCurrency(""),
    payments.WithPaymentsMethodsCreateFeeType(""),
    payments.WithPaymentsMethodsCreateKind(""),
    payments.WithPaymentsMethodsCreateLabels(map[string]interface{}{}),
    payments.WithPaymentsMethodsCreateMaxOrderValue(0),
    payments.WithPaymentsMethodsCreateMetadata(map[string]interface{}{}),
    payments.WithPaymentsMethodsCreateMinOrderValue(0),
    payments.WithPaymentsMethodsCreatePosition(0),
    payments.WithPaymentsMethodsCreateProvider(""),
    payments.WithPaymentsMethodsCreateProviderMethod(""),
)
```
