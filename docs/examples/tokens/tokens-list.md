```go
package main

import (
    "fmt"
    "github.com/revenexx-sdks/go/client"
    "github.com/revenexx-sdks/go/tokens"
)

client := client.New(
    client.WithEndpoint("https://api.revenexx.com")
    client.WithApiKeyAuth("<API_KEY>")
)

service := tokens.New(client)

response, error := service.TokensList(
    "",
    "",
    tokens.WithTokensListQueries([]interface{}{}),
    tokens.WithTokensListTotal(false),
)
```
