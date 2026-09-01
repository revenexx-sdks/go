package models

import (
    "encoding/json"
    "errors"
)

// AuthRecoveryResponse The identity service's recovery token, minus its
// secret, plus which mail the customer got. The secret is stripped
// deliberately — it travels only in the mailed link, and a caller that had
// both would not need the mail at all. `mail` is `tenant` when this shop's
// own template went out and `platform` when the messaging service could not
// be reached and the identity service's built-in mail is the copy the buyer
// has; the link is the same either way.
type AuthRecoveryResponse struct {
    // The recovery that was created.
    Id string `json:"$id"`
    // When the link stops working. The mail says the same thing in words.
    Expire string `json:"expire"`
    // Which template the buyer received: 'tenant' is this shop's own, 'platform'
    // the identity service's built-in one — the fallback when messaging could
    // not be reached. The link is identical either way, so a reset works in both
    // cases.
    Mail string `json:"mail"`
    // The platform user it belongs to.
    UserId string `json:"userId"`

    // Used by Decode() method
    data []byte
}

func (model AuthRecoveryResponse) New(data []byte) *AuthRecoveryResponse {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *AuthRecoveryResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}