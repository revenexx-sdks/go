package models

import (
    "encoding/json"
    "errors"
)

// AuthMfaChallengeConfirmResponse The identity service's answer, forwarded
// verbatim.
type AuthMfaChallengeConfirmResponse struct {
    // The challenge that was answered.
    Id string `json:"$id"`

    // Used by Decode() method
    data []byte
}

func (model AuthMfaChallengeConfirmResponse) New(data []byte) *AuthMfaChallengeConfirmResponse {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *AuthMfaChallengeConfirmResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}