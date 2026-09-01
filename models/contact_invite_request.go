package models

import (
    "encoding/json"
    "errors"
)

// ContactInviteRequest model.
type ContactInviteRequest struct {
    // Who did the inviting, as the recipient should read it. Absent, the company
    // name is used — "Beispiel GmbH invited you" reads better than the name of
    // somebody they have never heard of.
    InvitedBy string `json:"invited_by"`
    // Where the invitation points — the storefront sign-in, normally. There is
    // no token in it: the person is already a member and only has to sign in.
    Url string `json:"url"`

    // Used by Decode() method
    data []byte
}

func (model ContactInviteRequest) New(data []byte) *ContactInviteRequest {
    model.data = data
    return &model
}

func (model *ContactInviteRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}