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

response, error := service.FormsSubmissionsUpdate(
    "",
    forms.WithFormsSubmissionsUpdateData(map[string]interface{}{
        "company": "Example GmbH",
        "email": "buyer@example.com",
        "message": "Please quote 200 units of ACME-4711-BLK, delivered to Hamburg."
    }),
    forms.WithFormsSubmissionsUpdateFormId(""),
    forms.WithFormsSubmissionsUpdateFormSlug("contact"),
    forms.WithFormsSubmissionsUpdateMetadata(map[string]interface{}{}),
    forms.WithFormsSubmissionsUpdateSource("/contact"),
    forms.WithFormsSubmissionsUpdateStatus("new"),
)
```
