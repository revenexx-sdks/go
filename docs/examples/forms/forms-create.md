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

response, error := service.FormsCreate(
    "Price request",
    "price-request",
    forms.WithFormsCreateDefinition(interface{}{{"$formkit":"text","label":"Company","name":"company","validation":"required"},{"$formkit":"email","label":"Email","name":"email","validation":"required|email"},{"$formkit":"textarea","label":"What do you need a price for?","name":"message","rows":4},{"$el":"p","children":"We answer price requests within one working day."}}),
    forms.WithFormsCreateMetadata(map[string]interface{}{}),
    forms.WithFormsCreateSettings(map[string]interface{}{}),
    forms.WithFormsCreateStatus("draft"),
)
```
