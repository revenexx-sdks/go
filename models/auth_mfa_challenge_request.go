package models

import (
    "encoding/json"
    "errors"
)

// AuthMfaChallengeRequest model.
type AuthMfaChallengeRequest struct {
    // Which factor to challenge. Defaults to `email`, the only one this route
    // mails.
    Factor string `json:"factor"`
    // The platform user being challenged.
    UserId string `json:"user_id"`

    // Used by Decode() method
    data []byte
}

func (model AuthMfaChallengeRequest) New(data []byte) *AuthMfaChallengeRequest {
    model.data = data
    return &model
}

func (model *AuthMfaChallengeRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}