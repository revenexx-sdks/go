package models

import (
    "encoding/json"
    "errors"
)

// ProductAssociationsCreateRequest model.
type ProductAssociationsCreateRequest struct {
    // Which kind of relation this is — the `association_types` row.
    AssociationTypeId string `json:"association_type_id"`
    // Order in which the targets are shown, ascending.
    Position int `json:"position"`
    // The product the relation starts at — the one whose detail page shows it.
    ProductId string `json:"product_id"`
    // How many of the target belong to the source — the 4 in "this bundle
    // contains 4 casters". Only meaningful when the association type carries
    // `is_quantified`; null on an ordinary cross-sell.
    Quantity float64 `json:"quantity"`
    // The product the relation points at — the accessory, the spare part, the
    // cross-sell.
    TargetProductId string `json:"target_product_id"`

    // Used by Decode() method
    data []byte
}

func (model ProductAssociationsCreateRequest) New(data []byte) *ProductAssociationsCreateRequest {
    model.data = data
    return &model
}

func (model *ProductAssociationsCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}