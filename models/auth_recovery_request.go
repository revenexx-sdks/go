package models

import (
    "encoding/json"
    "errors"
)

// Model
type AuthRecoveryRequest struct {
    // 
    Email string `json:"email"`
    // Redirect URL carrying userId + secret.
    Url string `json:"url"`

    // Used by Decode() method
    data []byte
}

func (model AuthRecoveryRequest) New(data []byte) *AuthRecoveryRequest {
    model.data = data
    return &model
}

func (model *AuthRecoveryRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}