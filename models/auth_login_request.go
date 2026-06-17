package models

import (
    "encoding/json"
    "errors"
)

// Model
type AuthLoginRequest struct {
    // 
    Email string `json:"email"`
    // 
    Password string `json:"password"`

    // Used by Decode() method
    data []byte
}

func (model AuthLoginRequest) New(data []byte) *AuthLoginRequest {
    model.data = data
    return &model
}

func (model *AuthLoginRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}