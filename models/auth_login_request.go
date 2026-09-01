package models

import (
    "encoding/json"
    "errors"
)

// AuthLoginRequest model.
type AuthLoginRequest struct {
    // The buyer's login address — the same one the contact carries.
    Email string `json:"email"`
    // The password from registration or recovery. Wrong credentials are a 401; a
    // correct one on an undecided application is a 403.
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