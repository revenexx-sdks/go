package models

import (
    "encoding/json"
    "errors"
)

// AuthLoginResponse model.
type AuthLoginResponse struct {
    // The challenge to answer, when one was required. Send it back as
    // `challenge_id`.
    ChallengeId string `json:"challenge_id"`
    // The customer record behind the login. Null when a platform user has no
    // contact mirrored against it — a storefront should treat that as "signed
    // in, but not a customer of this app".
    Contact Contact `json:"contact"`
    // Present and true when the tenant's `mfa_mode` is 'required'. The password
    // was one of two things this buyer has to prove: a challenge has already been
    // created and mailed, and the session above must NOT be treated as signed in
    // until `PUT /customers/auth/mfa/challenge` confirms the code. The session
    // travels anyway because answering needs it — the expected caller holds
    // session material server-side, and this is the point at which that trust is
    // used.
    MfaRequired bool `json:"mfa_required"`
    // A contact's effective grants, derived from its role on every read —
    // nothing here is stored, so a role change can never leave a stale grant
    // behind. Carried here so a BFF does not need a second call to decide what to
    // render.
    Permissions ContactPermissions `json:"permissions"`
    // Platform auth session. Treat `secret` as a credential — the trusted BFF
    // stores it server-side (HTTP-only cookie), never in the browser.
    Session AuthSession `json:"session"`

    // Used by Decode() method
    data []byte
}

func (model AuthLoginResponse) New(data []byte) *AuthLoginResponse {
    model.data = data
    return &model
}

func (model *AuthLoginResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}