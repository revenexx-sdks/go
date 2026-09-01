package models

import (
    "encoding/json"
    "errors"
)

// ReferenceEntityRecordsFilter The exact-column filters this call was
// understood to carry, verbatim as they arrived. A query parameter that is
// not a column of `reference_entity_records` — `?status=`, a typo, a filter
// another entity has — is DROPPED and does not appear here, and the list
// comes back unfiltered. This object is the only way to tell that apart from
// "nothing matched".
type ReferenceEntityRecordsFilter struct {
    // The literal `?attribute_values=` value this call was understood to carry.
    AttributeValues string `json:"attribute_values"`
    // The literal `?code=` value this call was understood to carry.
    Code string `json:"code"`
    // The literal `?created_at=` value this call was understood to carry.
    CreatedAt string `json:"created_at"`
    // The literal `?id=` value this call was understood to carry.
    Id string `json:"id"`
    // The literal `?labels=` value this call was understood to carry.
    Labels string `json:"labels"`
    // The literal `?reference_entity_id=` value this call was understood to
    // carry.
    ReferenceEntityId string `json:"reference_entity_id"`
    // The literal `?updated_at=` value this call was understood to carry.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model ReferenceEntityRecordsFilter) New(data []byte) *ReferenceEntityRecordsFilter {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *ReferenceEntityRecordsFilter) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}