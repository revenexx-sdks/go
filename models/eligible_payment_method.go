package models

import (
    "encoding/json"
    "errors"
)

// Model
type EligiblePaymentMethod struct {
    // 
    Code string `json:"code"`
    // 
    Currency string `json:"currency"`
    // 
    Description string `json:"description"`
    // 
    Fee float64 `json:"fee"`
    // 
    FeeType string `json:"fee_type"`
    // 
    Kind string `json:"kind"`
    // 
    Labels interface{} `json:"labels"`
    // 
    Name string `json:"name"`
    // 
    Position int `json:"position"`
    // 
    Provider string `json:"provider"`

    // Used by Decode() method
    data []byte
}

func (model EligiblePaymentMethod) New(data []byte) *EligiblePaymentMethod {
    model.data = data
    return &model
}

func (model *EligiblePaymentMethod) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}