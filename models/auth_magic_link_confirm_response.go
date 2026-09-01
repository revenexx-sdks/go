package models

import (
    "encoding/json"
    "errors"
)

// AuthMagicLinkConfirmResponse model.
type AuthMagicLinkConfirmResponse struct {
    // The customer record behind the login. Null when no contact is mirrored
    // against the platform user yet — a sign-in link creates the account, not
    // the customer.
    Contact Contact `json:"contact"`
    // A contact's effective grants, derived from its role on every read —
    // nothing here is stored, so a role change can never leave a stale grant
    // behind. Null when there is no contact to derive them from.
    Permissions ContactPermissions `json:"permissions"`
    // Platform auth session. Treat `secret` as a credential — the trusted BFF
    // stores it server-side (HTTP-only cookie), never in the browser.
    Session AuthSession `json:"session"`

    // Used by Decode() method
    data []byte
}

func (model AuthMagicLinkConfirmResponse) New(data []byte) *AuthMagicLinkConfirmResponse {
    model.data = data
    return &model
}

func (model *AuthMagicLinkConfirmResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}