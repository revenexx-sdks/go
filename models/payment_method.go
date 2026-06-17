package models

import (
    "encoding/json"
    "errors"
)

// Model
type PaymentMethod struct {
    // 
    Code string `json:"code"`
    // 
    Countries interface{} `json:"countries"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Description string `json:"description"`
    // 
    Enabled bool `json:"enabled"`
    // 
    FeeAmount float64 `json:"fee_amount"`
    // 
    FeeCurrency string `json:"fee_currency"`
    // 
    FeeType string `json:"fee_type"`
    // 
    Id string `json:"id"`
    // 
    Kind string `json:"kind"`
    // 
    Labels interface{} `json:"labels"`
    // 
    MaxOrderValue float64 `json:"max_order_value"`
    // 
    Metadata interface{} `json:"metadata"`
    // 
    MinOrderValue float64 `json:"min_order_value"`
    // 
    Name string `json:"name"`
    // 
    Position int `json:"position"`
    // 
    Provider string `json:"provider"`
    // 
    ProviderMethod string `json:"provider_method"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model PaymentMethod) New(data []byte) *PaymentMethod {
    model.data = data
    return &model
}

func (model *PaymentMethod) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}