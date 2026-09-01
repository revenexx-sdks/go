```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/orders"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := orders.New(client)

response, error := service.OrdersList(
    orders.WithOrdersListId(""),
    orders.WithOrdersListNumber("ORD-000123"),
    orders.WithOrdersListCustomerOrderNumber("PO-2026-0042"),
    orders.WithOrdersListExternalRef("ERP-4711"),
    orders.WithOrdersListAcknowledgedAt("2026-01-01T12:00:00Z"),
    orders.WithOrdersListCartId(""),
    orders.WithOrdersListContactId(""),
    orders.WithOrdersListOrganizationId(""),
    orders.WithOrdersListChannelId(""),
    orders.WithOrdersListCurrency("EUR"),
    orders.WithOrdersListStatus("pending"),
    orders.WithOrdersListPaymentStatus("open"),
    orders.WithOrdersListFulfillmentStatus("unfulfilled"),
    orders.WithOrdersListOnHold(true),
    orders.WithOrdersListHoldReason("Credit check pending"),
    orders.WithOrdersListItemCount(3),
    orders.WithOrdersListSubtotal(149.7),
    orders.WithOrdersListShippingTotal(5.9),
    orders.WithOrdersListTaxTotal(29.56),
    orders.WithOrdersListGrandTotal(185.16),
    orders.WithOrdersListPlacedAt("2026-01-01T12:00:00Z"),
    orders.WithOrdersListCompletedAt("2026-01-01T12:00:00Z"),
    orders.WithOrdersListCancelledAt("2026-01-01T12:00:00Z"),
    orders.WithOrdersListCreatedAt("2026-01-01T12:00:00Z"),
    orders.WithOrdersListUpdatedAt("2026-01-01T12:00:00Z"),
    orders.WithOrdersListLimit(50),
    orders.WithOrdersListOffset(0),
    orders.WithOrdersListOrder("created_at.desc"),
)
```
