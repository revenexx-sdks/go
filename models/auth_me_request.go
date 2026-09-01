package models

import (
    "encoding/json"
    "errors"
)

// AuthMeRequest model.
type AuthMeRequest struct {
    // Optional session to verify. Pass it to ask "is this session still alive?"
    // (a revoked one is then a 401); omit it to only ask who a user is.
    SessionId string `json:"session_id"`
    // The platform user to resolve — `session.userId` from the login.
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