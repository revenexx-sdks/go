package models

import (
    "encoding/json"
    "errors"
)

// AuthRecoveryRequest model.
type AuthRecoveryRequest struct {
    // Who to send the recovery mail to. An address nobody holds is not
    // distinguished here — do not build an account-existence check on the
    // answer.
    Email string `json:"email"`
    // Where the mailed link points. `userId`, `secret` and `expire` are appended
    // as query parameters — the first two are what the confirm call takes. Same
    // shape the identity service's own mail used, so a storefront that already
    // handles that link needs no change.
    Url string `json:"url"`

    // Used by Decode() method
    data []byte
}

func (model AuthRecoveryRequest) New(data []byte) *AuthRecoveryRequest {
    model.data = data
    return &model
}

func (model *AuthRecoveryRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}