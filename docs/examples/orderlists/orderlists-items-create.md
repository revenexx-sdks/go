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

response, error := service.OrderlistsItemsCreate(
    "",
    "Copy paper A4, 80 g/m², white",
    orderlists.WithOrderlistsItemsCreateCategorySlug("office-supplies"),
    orderlists.WithOrderlistsItemsCreateCostCenterId("CC-100"),
    orderlists.WithOrderlistsItemsCreateCustomSku("CUST-4711"),
    orderlists.WithOrderlistsItemsCreateImage("https://cdn.example.com/catalog/acme-4711-blk.jpg"),
    orderlists.WithOrderlistsItemsCreateMetadata(map[string]interface{}{
        "erp_line_ref": "4711-01"
    }),
    orderlists.WithOrderlistsItemsCreatePosition(0),
    orderlists.WithOrderlistsItemsCreatePositionTexts(interface{}{"Deliver to bay 3","Engraving: Team A"}),
    orderlists.WithOrderlistsItemsCreatePrice(3.49),
    orderlists.WithOrderlistsItemsCreateProductId(""),
    orderlists.WithOrderlistsItemsCreateQuantity(12),
    orderlists.WithOrderlistsItemsCreateSku("ACME-4711-BLK"),
    orderlists.WithOrderlistsItemsCreateSubcategorySlug("paper"),
    orderlists.WithOrderlistsItemsCreateTaxRate(19),
    orderlists.WithOrderlistsItemsCreateUnit("piece"),
)
```
