package models

import (
    "encoding/json"
    "errors"
)

// Model
type AuthRecoveryConfirmRequest struct {
    // 
    Password string `json:"password"`
    // 
    Secret string `json:"secret"`
    // 
    UserId string `json:"user_id"`

    // Used by Decode() method
    data []byte
}

func (model AuthRecoveryConfirmRequest) New(data []byte) *AuthRecoveryConfirmRequest {
    model.data = data
    return &model
}

func (model *AuthRecoveryConfirmRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}