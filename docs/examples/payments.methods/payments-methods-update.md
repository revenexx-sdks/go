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

response, error := service.PaymentsMethodsUpdate(
    "",
    payments_methods.WithPaymentsMethodsUpdateCode("invoice"),
    payments_methods.WithPaymentsMethodsUpdateCountries(interface{}{"DE","AT"}),
    payments_methods.WithPaymentsMethodsUpdateDescription("Pay within 14 days of the invoice date."),
    payments_methods.WithPaymentsMethodsUpdateEnabled(true),
    payments_methods.WithPaymentsMethodsUpdateFeeAmount(2.5),
    payments_methods.WithPaymentsMethodsUpdateFeeCurrency("EUR"),
    payments_methods.WithPaymentsMethodsUpdateFeeType("none"),
    payments_methods.WithPaymentsMethodsUpdateKind("self_managed"),
    payments_methods.WithPaymentsMethodsUpdateLabels(map[string]interface{}{
        "de": "Rechnung",
        "en": "Invoice"
    }),
    payments_methods.WithPaymentsMethodsUpdateMaxOrderValue(2500),
    payments_methods.WithPaymentsMethodsUpdateMetadata(map[string]interface{}{
        "erp_payment_key": "ZTRM01"
    }),
    payments_methods.WithPaymentsMethodsUpdateMinOrderValue(10),
    payments_methods.WithPaymentsMethodsUpdateName("Invoice"),
    payments_methods.WithPaymentsMethodsUpdatePosition(0),
    payments_methods.WithPaymentsMethodsUpdateProvider("stripe"),
    payments_methods.WithPaymentsMethodsUpdateProviderMethod("card"),
)
```
