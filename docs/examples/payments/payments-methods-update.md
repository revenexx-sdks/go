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

response, error := service.PaymentsMethodsUpdate(
    "",
    payments.WithPaymentsMethodsUpdateCode(""),
    payments.WithPaymentsMethodsUpdateCountries([]interface{}{}),
    payments.WithPaymentsMethodsUpdateDescription(""),
    payments.WithPaymentsMethodsUpdateEnabled(false),
    payments.WithPaymentsMethodsUpdateFeeAmount(0),
    payments.WithPaymentsMethodsUpdateFeeCurrency(""),
    payments.WithPaymentsMethodsUpdateFeeType(""),
    payments.WithPaymentsMethodsUpdateKind(""),
    payments.WithPaymentsMethodsUpdateLabels(map[string]interface{}{}),
    payments.WithPaymentsMethodsUpdateMaxOrderValue(0),
    payments.WithPaymentsMethodsUpdateMetadata(map[string]interface{}{}),
    payments.WithPaymentsMethodsUpdateMinOrderValue(0),
    payments.WithPaymentsMethodsUpdateName(""),
    payments.WithPaymentsMethodsUpdatePosition(0),
    payments.WithPaymentsMethodsUpdateProvider(""),
    payments.WithPaymentsMethodsUpdateProviderMethod(""),
)
```
