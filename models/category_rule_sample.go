package models

import (
    "encoding/json"
    "errors"
)

// CategoryRuleSample model.
type CategoryRuleSample struct {
    // A matching product.
    Id string `json:"id"`
    // Its SKU, so the sample is readable. Null only for a row whose SKU is unset,
    // which the database does not allow.
    Sku string `json:"sku"`

    // Used by Decode() method
    data []byte
}

func (model CategoryRuleSample) New(data []byte) *CategoryRuleSample {
    model.data = data
    return &model
}

func (model *CategoryRuleSample) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}