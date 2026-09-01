package models

import (
    "encoding/json"
    "errors"
)

// AuthMagicLinkRequest model.
type AuthMagicLinkRequest struct {
    // Who to send the link to. An address that has never been seen creates an
    // account rather than failing.
    Email string `json:"email"`
    // Where the mailed link points. `userId`, `secret` and `expire` are appended
    // as query parameters; the first two are what the confirm call takes.
    Url string `json:"url"`

    // Used by Decode() method
    data []byte
}

func (model AuthMagicLinkRequest) New(data []byte) *AuthMagicLinkRequest {
    model.data = data
    return &model
}

func (model *AuthMagicLinkRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}