package models

import (
    "encoding/json"
    "errors"
)

// ProductCompleteness What was measured and stored into
// `products.completeness` by this call — how much of what the family
// requires the product actually carries.
type ProductCompleteness struct {
    // When this measurement was taken. It is a snapshot: editing the product does
    // not update it, the next `POST /products/{id}/completeness` does.
    ComputedAt string `json:"computed_at"`
    // How many of those carry a value — in ANY bucket, so a name held only in
    // German counts.
    Filled int `json:"filled"`
    // Attribute codes with no value in any bucket.
    Missing []string `json:"missing"`
    // filled / required, 0..1. A family that requires nothing is 1, not
    // undefined.
    Ratio float64 `json:"ratio"`
    // Attributes the product's family marks is_required.
    Required int `json:"required"`

    // Used by Decode() method
    data []byte
}

func (model ProductCompleteness) New(data []byte) *ProductCompleteness {
    model.data = data
    return &model
}

func (model *ProductCompleteness) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}