```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/orderlists"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := orderlists.New(client)

response, error := service.OrderlistsItemsList(
    "",
    orderlists.WithOrderlistsItemsListId(""),
    orderlists.WithOrderlistsItemsListProductId(""),
    orderlists.WithOrderlistsItemsListSku("ACME-4711-BLK"),
    orderlists.WithOrderlistsItemsListName("Copy paper A4, 80 g/m², white"),
    orderlists.WithOrderlistsItemsListImage("https://cdn.example.com/catalog/acme-4711-blk.jpg"),
    orderlists.WithOrderlistsItemsListQuantity(12),
    orderlists.WithOrderlistsItemsListUnit("piece"),
    orderlists.WithOrderlistsItemsListPrice(3.49),
    orderlists.WithOrderlistsItemsListTaxRate(19),
    orderlists.WithOrderlistsItemsListCostCenterId("CC-100"),
    orderlists.WithOrderlistsItemsListPositionTexts("{}"),
    orderlists.WithOrderlistsItemsListCustomSku("CUST-4711"),
    orderlists.WithOrderlistsItemsListCategorySlug("office-supplies"),
    orderlists.WithOrderlistsItemsListSubcategorySlug("paper"),
    orderlists.WithOrderlistsItemsListPosition(0),
    orderlists.WithOrderlistsItemsListMetadata("{}"),
    orderlists.WithOrderlistsItemsListCreatedAt("2026-01-01T12:00:00Z"),
    orderlists.WithOrderlistsItemsListUpdatedAt("2026-01-01T12:00:00Z"),
    orderlists.WithOrderlistsItemsListLimit(50),
    orderlists.WithOrderlistsItemsListOffset(0),
    orderlists.WithOrderlistsItemsListOrder("created_at.desc"),
)
```
