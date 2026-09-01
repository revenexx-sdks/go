package models

import (
    "encoding/json"
    "errors"
)

// MarketFilter The exact-column filters this call applied, echoed back. Every
// value is the raw query string, never the column's own type:
// `?is_default=true` comes back as `"true"`. A `?column=value` naming a
// column this entity does not have is DROPPED rather than refused — the
// call answers 200 with the unfiltered list, and the key missing from here is
// the only way to find out.
type MarketFilter struct {
    // The `code` filter as it arrived, verbatim. Present only when the call sent
    // it.
    Code string `json:"code"`
    // The `created_at` filter as it arrived, verbatim. Present only when the call
    // sent it. Any form the database accepts as a timestamp, including a bare
    // date.
    CreatedAt string `json:"created_at"`
    // The `currency` filter as it arrived, verbatim. Present only when the call
    // sent it.
    Currency string `json:"currency"`
    // The `id` filter as it arrived, verbatim. Present only when the call sent
    // it.
    Id string `json:"id"`
    // The `is_default` filter as it arrived, verbatim. Present only when the call
    // sent it.
    IsDefault string `json:"is_default"`
    // The `labels` filter as it arrived, verbatim. Present only when the call
    // sent it.
    Labels string `json:"labels"`
    // The `name` filter as it arrived, verbatim. Present only when the call sent
    // it.
    Name string `json:"name"`
    // The `position` filter as it arrived, verbatim. Present only when the call
    // sent it.
    Position string `json:"position"`
    // The `status` filter as it arrived, verbatim. Present only when the call
    // sent it.
    Status string `json:"status"`
    // The `updated_at` filter as it arrived, verbatim. Present only when the call
    // sent it. Any form the database accepts as a timestamp, including a bare
    // date.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model MarketFilter) New(data []byte) *MarketFilter {
    model.data = data
    return &model
}

func (model *MarketFilter) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}