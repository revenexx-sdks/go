package models

import (
    "encoding/json"
    "errors"
)

// Language Model
type Language struct {
    // Language two-character ISO 639-1 codes.
    Code string `json:"code"`
    // Language name.
    Name string `json:"name"`
    // Language native name.
    NativeName string `json:"nativeName"`

    // Used by Decode() method
    data []byte
}

func (model Language) New(data []byte) *Language {
    model.data = data
    return &model
}

func (model *Language) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}