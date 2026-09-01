package models

import (
    "encoding/json"
    "errors"
)

// PageTranslateRequest The strings to translate. They are forwarded to the
// tenant's provider verbatim.
type PageTranslateRequest struct {
    // The strings to translate. This app reads no element of the list — the
    // provider defines the contract, and the blökkli adapter sends the fields
    // below.
    Items []interface{} `json:"items"`

    // Used by Decode() method
    data []byte
}

func (model PageTranslateRequest) New(data []byte) *PageTranslateRequest {
    model.data = data
    return &model
}

func (model *PageTranslateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}