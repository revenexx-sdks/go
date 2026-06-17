package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type MarketUpdateRequest struct {
    // Market code (unique per tenant).
    Code string `json:"code"`
    // ISO 4217 code (default 'EUR').
    Currency string `json:"currency"`
    // 
    IsDefault bool `json:"is_default"`
    // Localized display names ({locale: label}).
    Labels interface{} `json:"labels"`
    // 
    Name string `json:"name"`
    // Sort position (default 0).
    Position int `json:"position"`
    // Default 'active'.
    Status string `json:"status"`

    // Used by Decode() method
    data []byte
}

func (model MarketUpdateRequest) New(data []byte) *MarketUpdateRequest {
    model.data = data
    return &model
}

func (model *MarketUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}