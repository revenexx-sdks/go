package models

import (
    "encoding/json"
    "errors"
)

// ReferenceEntitiesUpdateRequest Partial update — omitted fields keep their
// current value.
type ReferenceEntitiesUpdateRequest struct {
    // The entity's stable identifier — a domain of records the catalog POINTS
    // AT instead of duplicating, so a brand is edited once and not on nine
    // thousand products. Unique per tenant.
    Code string `json:"code"`
    // A delivery path or URL for the entity's own icon. Cosmetic — nothing in
    // this app resolves it.
    Image string `json:"image"`
    // What the entity is called, per language tag — the heading over its record
    // list.
    Labels interface{} `json:"labels"`

    // Used by Decode() method
    data []byte
}

func (model ReferenceEntitiesUpdateRequest) New(data []byte) *ReferenceEntitiesUpdateRequest {
    model.data = data
    return &model
}

func (model *ReferenceEntitiesUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}