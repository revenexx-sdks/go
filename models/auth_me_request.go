package models

import (
    "encoding/json"
    "errors"
)

// Model
type AuthMeRequest struct {
    // Optional session to verify — answers 401 when the session is expired or
    // revoked.
    SessionId string `json:"session_id"`
    // 
    UserId string `json:"user_id"`

    // Used by Decode() method
    data []byte
}

func (model AuthMeRequest) New(data []byte) *AuthMeRequest {
    model.data = data
    return &model
}

func (model *AuthMeRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}