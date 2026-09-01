package models

import (
    "encoding/json"
    "errors"
)

// AuthVerificationRequest model.
type AuthVerificationRequest struct {
    // Where the mailed link points. `userId`, `secret` and `expire` are appended
    // as query parameters; the first two are what the confirm call takes.
    Url string `json:"url"`
    // The platform user whose address is being confirmed — `user_id` from the
    // registration, or `session.userId` from a login.
    UserId string `json:"user_id"`

    // Used by Decode() method
    data []byte
}

func (model AuthVerificationRequest) New(data []byte) *AuthVerificationRequest {
    model.data = data
    return &model
}

func (model *AuthVerificationRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}