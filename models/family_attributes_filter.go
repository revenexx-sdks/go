package models

import (
    "encoding/json"
    "errors"
)

// FamilyAttributesFilter The exact-column filters this call was understood to
// carry, verbatim as they arrived. A query parameter that is not a column of
// `family_attributes` — `?status=`, a typo, a filter another entity has —
// is DROPPED and does not appear here, and the list comes back unfiltered.
// This object is the only way to tell that apart from "nothing matched".
type FamilyAttributesFilter struct {
    // The literal `?attribute_id=` value this call was understood to carry.
    AttributeId string `json:"attribute_id"`
    // The literal `?created_at=` value this call was understood to carry.
    CreatedAt string `json:"created_at"`
    // The literal `?family_id=` value this call was understood to carry.
    FamilyId string `json:"family_id"`
    // The literal `?id=` value this call was understood to carry.
    Id string `json:"id"`
    // The literal `?is_required=` value this call was understood to carry.
    IsRequired string `json:"is_required"`
    // The literal `?position=` value this call was understood to carry.
    Position string `json:"position"`
    // The literal `?required_channels=` value this call was understood to carry.
    RequiredChannels string `json:"required_channels"`

    // Used by Decode() method
    data []byte
}

func (model FamilyAttributesFilter) New(data []byte) *FamilyAttributesFilter {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *FamilyAttributesFilter) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}