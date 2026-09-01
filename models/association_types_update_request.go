package models

import (
    "encoding/json"
    "errors"
)

// AssociationTypesUpdateRequest Partial update — omitted fields keep their
// current value.
type AssociationTypesUpdateRequest struct {
    // The kind of relation between two products. Unique per tenant.
    Code string `json:"code"`
    // Declares that a relation of this kind carries a quantity — a bundle, a
    // bill of materials. `product_associations.quantity` is where that number
    // goes, and it is meaningless without this flag.
    IsQuantified bool `json:"is_quantified"`
    // Declares the relation symmetric — an accessory of A is an accessory of B.
    // It is a declaration a client reads: this app stores one row per direction
    // and does not create the mirror for you.
    IsTwoWay bool `json:"is_two_way"`
    // What the relation is called in a product form, per language tag.
    Labels interface{} `json:"labels"`

    // Used by Decode() method
    data []byte
}

func (model AssociationTypesUpdateRequest) New(data []byte) *AssociationTypesUpdateRequest {
    model.data = data
    return &model
}

func (model *AssociationTypesUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}