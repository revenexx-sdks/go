```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/forms"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com"),
    client.WithTenant("<TENANT_SLUG>"),
    client.WithApiKeyAuth("<API_KEY>"),
)

service := forms.New(client)

response, error := service.FormsSubmissionsCreate(
    map[string]interface{}{
        "company": "Example GmbH",
        "email": "buyer@example.com",
        "message": "Please quote 200 units of ACME-4711-BLK, delivered to Hamburg."
    },
    "",
    forms.WithFormsSubmissionsCreateFormSlug("contact"),
    forms.WithFormsSubmissionsCreateMetadata(map[string]interface{}{}),
    forms.WithFormsSubmissionsCreateSource("/contact"),
    forms.WithFormsSubmissionsCreateStatus("new"),
)
```
