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

response, error := service.PaymentsList(
    payments_ledger.WithPaymentsListLimit(1),
    payments_ledger.WithPaymentsListOffset(1),
    payments_ledger.WithPaymentsListOrder("created_at.desc"),
    payments_ledger.WithPaymentsListCartId(""),
    payments_ledger.WithPaymentsListContactId(""),
    payments_ledger.WithPaymentsListStatus("created"),
    payments_ledger.WithPaymentsListOrderRef("ORD-10042"),
    payments_ledger.WithPaymentsListMethodCode("invoice"),
    payments_ledger.WithPaymentsListKind("self_managed"),
    payments_ledger.WithPaymentsListProvider("stripe"),
    payments_ledger.WithPaymentsListDunningStage("none"),
    payments_ledger.WithPaymentsListIdempotencyKey("checkout-2f9c41"),
)
```
