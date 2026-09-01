package models

import (
    "encoding/json"
    "errors"
)

// AuthRecoveryConfirmResponse The identity service's answer, forwarded
// verbatim: the spent recovery token. The new password is already in effect
// when this arrives.
type AuthRecoveryConfirmResponse struct {
    // The recovery that was confirmed.
    Id string `json:"$id"`
    // The platform user whose password was set.
    UserId string `json:"userId"`

    // Used by Decode() method
    data []byte
}

func (model AuthRecoveryConfirmResponse) New(data []byte) *AuthRecoveryConfirmResponse {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *AuthRecoveryConfirmResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}