package models

import (
    "encoding/json"
    "errors"
)

// AuthMfaChallengeConfirmRequest model.
type AuthMfaChallengeConfirmRequest struct {
    // The `$id` the send answered with.
    ChallengeId string `json:"challenge_id"`
    // What the buyer typed.
    Code string `json:"code"`
    // The same session the challenge was created with.
    SessionSecret string `json:"session_secret"`
    // The platform user, for the caller's own bookkeeping. The challenge already
    // knows whose it is.
    UserId string `json:"user_id"`

    // Used by Decode() method
    data []byte
}

func (model AuthMfaChallengeConfirmRequest) New(data []byte) *AuthMfaChallengeConfirmRequest {
    model.data = data
    return &model
}

func (model *AuthMfaChallengeConfirmRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}