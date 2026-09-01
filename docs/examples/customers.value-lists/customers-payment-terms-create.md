```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/customers_value_lists"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := customers_value_lists.New(client)

response, error := service.CustomersPaymentTermsCreate(
    "",
    "Net 30 days",
    customers_value_lists.WithCustomersPaymentTermsCreateDescription("Invoice due 30 days after the delivery note."),
    customers_value_lists.WithCustomersPaymentTermsCreateDescriptions(map[string]interface{}{
        "de": "Rechnung 30 Tage nach Lieferschein f\u00e4llig.",
        "en": "Invoice due 30 days after the delivery note."
    }),
    customers_value_lists.WithCustomersPaymentTermsCreateIsDefault(true),
    customers_value_lists.WithCustomersPaymentTermsCreateLabels(map[string]interface{}{
        "de": "Zahlbar in 30 Tagen",
        "en": "Net 30 days"
    }),
    customers_value_lists.WithCustomersPaymentTermsCreatePosition(1),
    customers_value_lists.WithCustomersPaymentTermsCreateTone("neutral"),
)
```
