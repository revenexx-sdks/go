package models

import (
    "encoding/json"
    "errors"
)

// AssociationTypesFilter The exact-column filters this call was understood to
// carry, verbatim as they arrived. A query parameter that is not a column of
// `association_types` — `?status=`, a typo, a filter another entity has —
// is DROPPED and does not appear here, and the list comes back unfiltered.
// This object is the only way to tell that apart from "nothing matched".
type AssociationTypesFilter struct {
    // The literal `?code=` value this call was understood to carry.
    Code string `json:"code"`
    // The literal `?created_at=` value this call was understood to carry.
    CreatedAt string `json:"created_at"`
    // The literal `?id=` value this call was understood to carry.
    Id string `json:"id"`
    // The literal `?is_quantified=` value this call was understood to carry.
    IsQuantified string `json:"is_quantified"`
    // The literal `?is_two_way=` value this call was understood to carry.
    IsTwoWay string `json:"is_two_way"`
    // The literal `?labels=` value this call was understood to carry.
    Labels string `json:"labels"`

    // Used by Decode() method
    data []byte
}

func (model AssociationTypesFilter) New(data []byte) *AssociationTypesFilter {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *AssociationTypesFilter) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}