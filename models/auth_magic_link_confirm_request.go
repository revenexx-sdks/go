package models

import (
    "encoding/json"
    "errors"
)

// AuthMagicLinkConfirmRequest model.
type AuthMagicLinkConfirmRequest struct {
    // The one-time secret the mailed link carried. Spent on first use and
    // expiring, so a second attempt with the same one is a 401 rather than a
    // second session.
    Secret string `json:"secret"`
    // The `userId` the mailed link carried.
    UserId string `json:"user_id"`

    // Used by Decode() method
    data []byte
}

func (model AuthMagicLinkConfirmRequest) New(data []byte) *AuthMagicLinkConfirmRequest {
    model.data = data
    return &model
}

func (model *AuthMagicLinkConfirmRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}