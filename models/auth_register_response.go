package models

import (
    "encoding/json"
    "errors"
)

// Model
type AuthRegisterResponse struct {
    // 
    Contact Contact `json:"contact"`
    // 
    UserId string `json:"user_id"`

    // Used by Decode() method
    data []byte
}

func (model AuthRegisterResponse) New(data []byte) *AuthRegisterResponse {
    model.data = data
    return &model
}

func (model *AuthRegisterResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}