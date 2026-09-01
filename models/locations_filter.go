package models

import (
    "encoding/json"
    "errors"
)

// LocationsFilter The exact-column filters this call was understood to carry,
// verbatim as they arrived. A query parameter that is not a column of
// `locations` — a typo, a filter another entity has, `?q=` — is DROPPED
// and cannot appear here, and the list comes back unfiltered. This object is
// the only way to tell that apart from "nothing matched".
type LocationsFilter struct {
    // The literal `?address=` value this call was understood to carry.
    Address string `json:"address"`
    // The literal `?code=` value this call was understood to carry.
    Code string `json:"code"`
    // The literal `?created_at=` value this call was understood to carry.
    CreatedAt string `json:"created_at"`
    // The literal `?enabled=` value this call was understood to carry.
    Enabled string `json:"enabled"`
    // The literal `?id=` value this call was understood to carry.
    Id string `json:"id"`
    // The literal `?labels=` value this call was understood to carry.
    Labels string `json:"labels"`
    // The literal `?metadata=` value this call was understood to carry.
    Metadata string `json:"metadata"`
    // The literal `?name=` value this call was understood to carry.
    Name string `json:"name"`
    // The literal `?priority=` value this call was understood to carry.
    Priority string `json:"priority"`
    // The literal `?type=` value this call was understood to carry.
    Type string `json:"type"`
    // The literal `?updated_at=` value this call was understood to carry.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model LocationsFilter) New(data []byte) *LocationsFilter {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *LocationsFilter) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}