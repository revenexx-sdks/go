package models

import (
    "encoding/json"
    "errors"
)

// Model
type ProductTaxRef struct {
    // 
    Id string `json:"id"`
    // 
    Sku string `json:"sku"`
    // 
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