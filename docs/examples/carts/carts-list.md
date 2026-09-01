```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/carts"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := carts.New(client)

response, error := service.CartsList(
    carts.WithCartsListId(""),
    carts.WithCartsListName("Weekly order"),
    carts.WithCartsListStatus("active"),
    carts.WithCartsListContactId(""),
    carts.WithCartsListSessionKey("a1b2c3d4e5f6"),
    carts.WithCartsListChannelId(""),
    carts.WithCartsListCurrency("EUR"),
    carts.WithCartsListIsCurrent(true),
    carts.WithCartsListItemCount(100),
    carts.WithCartsListSubtotal(12),
    carts.WithCartsListAbandonedAt("2026-01-01T12:00:00Z"),
    carts.WithCartsListOrderedAt("2026-01-01T12:00:00Z"),
    carts.WithCartsListOrderRef("SO-10042"),
    carts.WithCartsListMergedIntoCartId(""),
    carts.WithCartsListCreatedAt("2026-01-01T12:00:00Z"),
    carts.WithCartsListUpdatedAt("2026-01-01T12:00:00Z"),
    carts.WithCartsListLimit(1),
    carts.WithCartsListOffset(1),
    carts.WithCartsListOrder("created_at.desc"),
)
```
