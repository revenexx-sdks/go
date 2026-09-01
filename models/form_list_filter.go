package models

import (
    "encoding/json"
    "errors"
)

// FormListFilter The exact-column filters this call was understood to carry,
// echoed with the values as they arrived. A query parameter that is not a
// filterable column of this entity is DROPPED rather than refused, and is
// simply missing here — so an empty object next to a query string that had
// a filter in it means the filter was misspelled, and is the only way to tell
// that from a filter that matched nothing.
type FormListFilter struct {
    // The `created_at` filter, verbatim as the query string carried it. A string
    // here whatever the column's own type.
    CreatedAt string `json:"created_at"`
    // The `id` filter, verbatim as the query string carried it. A string here
    // whatever the column's own type.
    Id string `json:"id"`
    // The `name` filter, verbatim as the query string carried it. A string here
    // whatever the column's own type.
    Name string `json:"name"`
    // The `slug` filter, verbatim as the query string carried it. A string here
    // whatever the column's own type.
    Slug string `json:"slug"`
    // The `status` filter, verbatim as the query string carried it. A string here
    // whatever the column's own type — and NOT necessarily one of the permitted
    // values: `?status=zzz` is echoed back unchanged and matches nothing, which
    // is the point of the echo.
    Status string `json:"status"`
    // The `updated_at` filter, verbatim as the query string carried it. A string
    // here whatever the column's own type.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model FormListFilter) New(data []byte) *FormListFilter {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *FormListFilter) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}