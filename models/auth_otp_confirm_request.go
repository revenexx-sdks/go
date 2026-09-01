package models

import (
    "encoding/json"
    "errors"
)

// AuthOtpConfirmRequest model.
type AuthOtpConfirmRequest struct {
    // The one-time secret the mailed code carried. Spent on first use and
    // expiring, so a second attempt with the same one is a 401 rather than a
    // second session.
    Secret string `json:"secret"`
    // The `userId` the mailed code carried.
    UserId string `json:"user_id"`

    // Used by Decode() method
    data []byte
}

func (model AuthOtpConfirmRequest) New(data []byte) *AuthOtpConfirmRequest {
    model.data = data
    return &model
}

func (model *AuthOtpConfirmRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}