package models

import (
    "encoding/json"
    "errors"
)

// ProductAssociationsFilter The exact-column filters this call was understood
// to carry, verbatim as they arrived. A query parameter that is not a column
// of `product_associations` — `?status=`, a typo, a filter another entity
// has — is DROPPED and does not appear here, and the list comes back
// unfiltered. This object is the only way to tell that apart from "nothing
// matched".
type ProductAssociationsFilter struct {
    // The literal `?association_type_id=` value this call was understood to
    // carry.
    AssociationTypeId string `json:"association_type_id"`
    // The literal `?created_at=` value this call was understood to carry.
    CreatedAt string `json:"created_at"`
    // The literal `?id=` value this call was understood to carry.
    Id string `json:"id"`
    // The literal `?position=` value this call was understood to carry.
    Position string `json:"position"`
    // The literal `?product_id=` value this call was understood to carry.
    ProductId string `json:"product_id"`
    // The literal `?quantity=` value this call was understood to carry.
    Quantity string `json:"quantity"`
    // The literal `?target_product_id=` value this call was understood to carry.
    TargetProductId string `json:"target_product_id"`

    // Used by Decode() method
    data []byte
}

func (model ProductAssociationsFilter) New(data []byte) *ProductAssociationsFilter {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *ProductAssociationsFilter) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}