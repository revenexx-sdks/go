package models

import (
    "encoding/json"
    "errors"
)

// AuthVerificationConfirmResponse The identity service's answer, forwarded
// verbatim: the spent verification token.
type AuthVerificationConfirmResponse struct {
    // The verification that was confirmed.
    Id string `json:"$id"`
    // The platform user whose address is now confirmed.
    UserId string `json:"userId"`

    // Used by Decode() method
    data []byte
}

func (model AuthVerificationConfirmResponse) New(data []byte) *AuthVerificationConfirmResponse {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *AuthVerificationConfirmResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}