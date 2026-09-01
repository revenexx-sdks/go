```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/payments_ledger"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := payments_ledger.New(client)

response, error := service.PaymentsCreate(
    49.9,
    "invoice",
    payments_ledger.WithPaymentsCreateCartId(""),
    payments_ledger.WithPaymentsCreateContactId(""),
    payments_ledger.WithPaymentsCreateCountry("DE"),
    payments_ledger.WithPaymentsCreateCurrency("EUR"),
    payments_ledger.WithPaymentsCreateIdempotencyKey("checkout-2f9c41"),
    payments_ledger.WithPaymentsCreateMetadata(map[string]interface{}{
        "order_source": "web"
    }),
    payments_ledger.WithPaymentsCreateOrderRef("ORD-10042"),
    payments_ledger.WithPaymentsCreateReturnUrl("https://shop.example.com/checkout/return"),
)
```
