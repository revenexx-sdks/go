package models

import (
    "encoding/json"
    "errors"
)

// AuthOtpRequest model.
type AuthOtpRequest struct {
    // Who to send the code to. As with the sign-in link, an unknown address
    // creates an account rather than failing.
    Email string `json:"email"`

    // Used by Decode() method
    data []byte
}

func (model AuthOtpRequest) New(data []byte) *AuthOtpRequest {
    model.data = data
    return &model
}

func (model *AuthOtpRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}