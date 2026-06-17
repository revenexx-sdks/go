package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type MarketCurrencyUpdateRequest struct {
    // ISO 4217 code, e.g. EUR (unique per market).
    Code string `json:"code"`
    // 
    IsDefault bool `json:"is_default"`
    // Sort position (default 0).
    Position int `json:"position"`

    // Used by Decode() method
    data []byte
}

func (model MarketCurrencyUpdateRequest) New(data []byte) *MarketCurrencyUpdateRequest {
    model.data = data
    return &model
}

func (model *MarketCurrencyUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}