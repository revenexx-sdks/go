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

response, error := service.PaymentsMethodsList(
    payments_methods.WithPaymentsMethodsListLimit(1),
    payments_methods.WithPaymentsMethodsListOffset(1),
    payments_methods.WithPaymentsMethodsListOrder("created_at.desc"),
    payments_methods.WithPaymentsMethodsListCode("invoice"),
    payments_methods.WithPaymentsMethodsListKind("self_managed"),
    payments_methods.WithPaymentsMethodsListEnabled(true),
    payments_methods.WithPaymentsMethodsListProvider("stripe"),
)
```
