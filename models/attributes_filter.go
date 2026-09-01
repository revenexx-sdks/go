package models

import (
    "encoding/json"
    "errors"
)

// AttributesFilter The exact-column filters this call was understood to
// carry, verbatim as they arrived. A query parameter that is not a column of
// `attributes` — `?status=`, a typo, a filter another entity has — is
// DROPPED and does not appear here, and the list comes back unfiltered. This
// object is the only way to tell that apart from "nothing matched".
type AttributesFilter struct {
    // The literal `?code=` value this call was understood to carry.
    Code string `json:"code"`
    // The literal `?config=` value this call was understood to carry.
    Config string `json:"config"`
    // The literal `?created_at=` value this call was understood to carry.
    CreatedAt string `json:"created_at"`
    // The literal `?entity_ref=` value this call was understood to carry.
    EntityRef string `json:"entity_ref"`
    // The literal `?entity_type=` value this call was understood to carry.
    EntityType string `json:"entity_type"`
    // The literal `?group_id=` value this call was understood to carry.
    GroupId string `json:"group_id"`
    // The literal `?id=` value this call was understood to carry.
    Id string `json:"id"`
    // The literal `?is_filterable=` value this call was understood to carry.
    IsFilterable string `json:"is_filterable"`
    // The literal `?is_unique=` value this call was understood to carry.
    IsUnique string `json:"is_unique"`
    // The literal `?labels=` value this call was understood to carry.
    Labels string `json:"labels"`
    // The literal `?localizable=` value this call was understood to carry.
    Localizable string `json:"localizable"`
    // The literal `?position=` value this call was understood to carry.
    Position string `json:"position"`
    // The literal `?scopable=` value this call was understood to carry.
    Scopable string `json:"scopable"`
    // The literal `?type=` value this call was understood to carry.
    Type string `json:"type"`
    // The literal `?updated_at=` value this call was understood to carry.
    UpdatedAt string `json:"updated_at"`
    // The literal `?usable_in_grid=` value this call was understood to carry.
    UsableInGrid string `json:"usable_in_grid"`
    // The literal `?validation=` value this call was understood to carry.
    Validation string `json:"validation"`

    // Used by Decode() method
    data []byte
}

func (model AttributesFilter) New(data []byte) *AttributesFilter {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *AttributesFilter) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}