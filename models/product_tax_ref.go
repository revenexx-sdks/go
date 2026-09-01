package models

import (
    "encoding/json"
    "errors"
)

// ProductTaxRef model.
type ProductTaxRef struct {
    // The product's id.
    Id string `json:"id"`
    // The product's resolved display name, or its SKU when the catalog holds no
    // name for it.
    Label string `json:"label"`
    // The SKU, so a caller that asked by id can key its own answer by SKU and the
    // other way round.
    Sku string `json:"sku"`
    // The tax class key the prices app resolves a rate from. Null means the
    // product names none and the caller has to fall back to its own default.
    TaxClass string `json:"tax_class"`

    // Used by Decode() method
    data []byte
}

func (model ProductTaxRef) New(data []byte) *ProductTaxRef {
    model.data = data
    return &model
}

func (model *ProductTaxRef) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}