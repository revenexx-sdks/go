package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type MarketTaxClassUpdateRequest struct {
    // Tax class code (unique per market).
    Code string `json:"code"`
    // 
    IsDefault bool `json:"is_default"`
    // Localized display names ({locale: label}).
    Labels interface{} `json:"labels"`
    // 
    Name string `json:"name"`
    // Sort position (default 0).
    Position int `json:"position"`
    // Tax rate in percent, 0–100 (default 0).
    Rate float64 `json:"rate"`

    // Used by Decode() method
    data []byte
}

func (model MarketTaxClassUpdateRequest) New(data []byte) *MarketTaxClassUpdateRequest {
    model.data = data
    return &model
}

func (model *MarketTaxClassUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}