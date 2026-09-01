package models

import (
    "encoding/json"
    "errors"
)

// MultiSearchEntry One search inside a federated request. `collection_name`
// is required on the gateway-trust path; with a `revx_` key it is optional
// and is forced to the key's own collection.
type MultiSearchEntry struct {
    // A collection the tenant owns.
    CollectionName string `json:"collection_name"`
    // Comma-separated document fields to omit.
    ExcludeFields string `json:"exclude_fields"`
    // Comma-separated fields to facet on.
    FacetBy string `json:"facet_by"`
    // Filter expression, e.g. `in_stock:=true && price:<100`. ANDed with the
    // tenant filter the proxy injects.
    FilterBy string `json:"filter_by"`
    // Comma-separated fields to group results by.
    GroupBy string `json:"group_by"`
    // Comma-separated fields to highlight in full.
    HighlightFullFields string `json:"highlight_full_fields"`
    // Comma-separated document fields to return.
    IncludeFields string `json:"include_fields"`
    // Facet values to return per field.
    MaxFacetValues int `json:"max_facet_values"`
    // Typos tolerated per query token.
    NumTypos int `json:"num_typos"`
    // 1-based page number.
    Page int `json:"page"`
    // Hits per page.
    PerPage int `json:"per_page"`
    // Whether the last token is a prefix; per-field when comma-separated.
    Prefix string `json:"prefix"`
    // Query text. Use `*` to match everything.
    Q string `json:"q"`
    // Comma-separated fields to search, in weight order.
    QueryBy string `json:"query_by"`
    // Sort expression, e.g. `price:desc`.
    SortBy string `json:"sort_by"`

    // Used by Decode() method
    data []byte
}

func (model MultiSearchEntry) New(data []byte) *MultiSearchEntry {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *MultiSearchEntry) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}