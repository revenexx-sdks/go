package models

import (
    "encoding/json"
    "errors"
)

// AuthMagicLinkResponse The token, minus its secret. The secret travels only
// in the mailed link — a caller holding both would not need the mail at
// all.
type AuthMagicLinkResponse struct {
    // The token that was created.
    Id string `json:"$id"`
    // When the link stops working. The mail says the same in words.
    Expire string `json:"expire"`
    // Which template the buyer received: 'tenant' is this shop's own, 'platform'
    // the identity service's built-in one — the fallback when messaging could
    // not be reached. The value is the same either way, so the flow works in both
    // cases.
    Mail string `json:"mail"`
    // The platform user it belongs to — new when the address was.
    UserId string `json:"userId"`

    // Used by Decode() method
    data []byte
}

func (model AuthMagicLinkResponse) New(data []byte) *AuthMagicLinkResponse {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *AuthMagicLinkResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}