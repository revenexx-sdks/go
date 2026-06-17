package models

import (
    "encoding/json"
    "errors"
)

// Model
type AuthLogoutRequest struct {
    // 
    SessionId string `json:"session_id"`
    // 
    UserId string `json:"user_id"`

    // Used by Decode() method
    data []byte
}

func (model AuthLogoutRequest) New(data []byte) *AuthLogoutRequest {
    model.data = data
    return &model
}

func (model *AuthLogoutRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}