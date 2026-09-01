```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/payments_providers"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := payments_providers.New(client)

response, error := service.PaymentsProvidersList(
    payments_providers.WithPaymentsProvidersListLimit(1),
    payments_providers.WithPaymentsProvidersListOffset(1),
    payments_providers.WithPaymentsProvidersListOrder("created_at.desc"),
    payments_providers.WithPaymentsProvidersListProvider("stripe"),
    payments_providers.WithPaymentsProvidersListEnabled(true),
    payments_providers.WithPaymentsProvidersListTestMode(true),
)
```
