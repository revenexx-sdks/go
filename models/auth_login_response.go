package models

import (
    "encoding/json"
    "errors"
)

// Model
type AuthLoginResponse struct {
    // 
    Contact Contact `json:"contact"`
    // 
    Session AuthSession `json:"session"`

    // Used by Decode() method
    data []byte
}

func (model AuthLoginResponse) New(data []byte) *AuthLoginResponse {
    model.data = data
    return &model
}

func (model *AuthLoginResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}