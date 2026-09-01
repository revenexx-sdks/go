package models

import (
    "encoding/json"
    "errors"
)

// AuthMeResponse model.
type AuthMeResponse struct {
    // The customer record mirrored against this user, or null. A user with no
    // contact resolves perfectly well — that is not the 404.
    Contact Contact `json:"contact"`
    // A contact's effective grants, derived from its role on every read —
    // nothing here is stored, so a role change can never leave a stale grant
    // behind. Null when there is no contact to derive them from.
    Permissions ContactPermissions `json:"permissions"`
    // The platform identity record, forwarded verbatim from the identity service.
    // This app neither reshapes nor validates it, so treat unknown fields as
    // forward-compatible; the ones named here are the ones this app itself writes
    // and reads.
    User interface{} `json:"user"`

    // Used by Decode() method
    data []byte
}

func (model AuthMeResponse) New(data []byte) *AuthMeResponse {
    model.data = data
    return &model
}

func (model *AuthMeResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}