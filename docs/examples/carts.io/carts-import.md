```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/carts_io"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := carts_io.New(client)

response, error := service.CartsImport(
    carts_io.WithCartsImportContactId(""),
    carts_io.WithCartsImportCsv("sku,name,quantity,unit_price
BOLT-M8-30,Hex bolt M8,100,0.12
NUT-M8,Hex nut M8,100,0.04
"),
    carts_io.WithCartsImportName("Weekly order"),
    carts_io.WithCartsImportPayload(map[string]interface{}{
        "cart": {
            "currency": "EUR",
            "name": "Weekly order"
        },
        "items": [
            {
                "name": "Hex bolt M8",
                "quantity": 100,
                "sku": "BOLT-M8-30",
                "unit_price": 0.12
            }
        ]
    }),
    carts_io.WithCartsImportProfileId(""),
    carts_io.WithCartsImportSessionKey("a1b2c3d4e5f6"),
    carts_io.WithCartsImportTargetCartId(""),
)
```
