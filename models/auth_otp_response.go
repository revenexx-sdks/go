package models

import (
    "encoding/json"
    "errors"
)

// AuthOtpResponse The token, minus the code. The code is in the mail and
// nowhere else.
type AuthOtpResponse struct {
    // The token that was created.
    Id string `json:"$id"`
    // When the code stops working.
    Expire string `json:"expire"`
    // Which template the buyer received: 'tenant' is this shop's own, 'platform'
    // the identity service's built-in one — the fallback when messaging could
    // not be reached. The value is the same either way, so the flow works in both
    // cases.
    Mail string `json:"mail"`
    // The platform user it belongs to — send it back with the code the buyer
    // types.
    UserId string `json:"userId"`

    // Used by Decode() method
    data []byte
}

func (model AuthOtpResponse) New(data []byte) *AuthOtpResponse {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *AuthOtpResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}