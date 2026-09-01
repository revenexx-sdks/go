package models

import (
    "encoding/json"
    "errors"
)

// ReferenceEntitiesCreateRequest model.
type ReferenceEntitiesCreateRequest struct {
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

func (model ReferenceEntitiesCreateRequest) New(data []byte) *ReferenceEntitiesCreateRequest {
    model.data = data
    return &model
}

func (model *ReferenceEntitiesCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}