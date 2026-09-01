package models

import (
    "encoding/json"
    "errors"
)

// AuthRecoveryConfirmRequest model.
type AuthRecoveryConfirmRequest struct {
    // The new password. It replaces the old one immediately; existing sessions
    // are the identity service's business, not this app's.
    Password string `json:"password"`
    // The one-time secret from the mailed link. Only that value works — it is
    // spent on first use and expires, and anything else is a 401, so no example
    // here would be anything but a call that fails.
    Secret string `json:"secret"`
    // The `userId` the mailed link carried.
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