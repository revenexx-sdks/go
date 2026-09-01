package models

import (
    "encoding/json"
    "errors"
)

// PaymentErrorRedactRequest model.
type PaymentErrorRedactRequest struct {
    // Write the reclassified values. Defaults to false, which reports what WOULD
    // change and touches nothing.
    Apply bool `json:"apply"`
    // How many payments to scan, oldest first. Defaults to 500, capped at 5000
    // — a tenant with more pre-taxonomy rows needs several runs, and re-running
    // is free.
    Limit int `json:"limit"`

    // Used by Decode() method
    data []byte
}

func (model PaymentErrorRedactRequest) New(data []byte) *PaymentErrorRedactRequest {
    model.data = data
    return &model
}

func (model *PaymentErrorRedactRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}