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

response, error := service.PaymentsMethodsCreate(
    "invoice",
    "Invoice",
    payments_methods.WithPaymentsMethodsCreateCountries(interface{}{"DE","AT"}),
    payments_methods.WithPaymentsMethodsCreateDescription("Pay within 14 days of the invoice date."),
    payments_methods.WithPaymentsMethodsCreateEnabled(true),
    payments_methods.WithPaymentsMethodsCreateFeeAmount(2.5),
    payments_methods.WithPaymentsMethodsCreateFeeCurrency("EUR"),
    payments_methods.WithPaymentsMethodsCreateFeeType("none"),
    payments_methods.WithPaymentsMethodsCreateKind("self_managed"),
    payments_methods.WithPaymentsMethodsCreateLabels(map[string]interface{}{
        "de": "Rechnung",
        "en": "Invoice"
    }),
    payments_methods.WithPaymentsMethodsCreateMaxOrderValue(2500),
    payments_methods.WithPaymentsMethodsCreateMetadata(map[string]interface{}{
        "erp_payment_key": "ZTRM01"
    }),
    payments_methods.WithPaymentsMethodsCreateMinOrderValue(10),
    payments_methods.WithPaymentsMethodsCreatePosition(0),
    payments_methods.WithPaymentsMethodsCreateProvider("stripe"),
    payments_methods.WithPaymentsMethodsCreateProviderMethod("card"),
)
```
