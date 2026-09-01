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

response, error := service.OrderlistsItemsUpdate(
    "",
    "",
    orderlists.WithOrderlistsItemsUpdateCategorySlug("office-supplies"),
    orderlists.WithOrderlistsItemsUpdateCostCenterId("CC-100"),
    orderlists.WithOrderlistsItemsUpdateCustomSku("CUST-4711"),
    orderlists.WithOrderlistsItemsUpdateImage("https://cdn.example.com/catalog/acme-4711-blk.jpg"),
    orderlists.WithOrderlistsItemsUpdateMetadata(map[string]interface{}{
        "erp_line_ref": "4711-01"
    }),
    orderlists.WithOrderlistsItemsUpdateName("Copy paper A4, 80 g/m², white"),
    orderlists.WithOrderlistsItemsUpdatePosition(0),
    orderlists.WithOrderlistsItemsUpdatePositionTexts(interface{}{"Deliver to bay 3","Engraving: Team A"}),
    orderlists.WithOrderlistsItemsUpdatePrice(3.49),
    orderlists.WithOrderlistsItemsUpdateProductId(""),
    orderlists.WithOrderlistsItemsUpdateQuantity(12),
    orderlists.WithOrderlistsItemsUpdateSku("ACME-4711-BLK"),
    orderlists.WithOrderlistsItemsUpdateSubcategorySlug("paper"),
    orderlists.WithOrderlistsItemsUpdateTaxRate(19),
    orderlists.WithOrderlistsItemsUpdateUnit("piece"),
)
```
