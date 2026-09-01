package models

import (
    "encoding/json"
    "errors"
)

// AuthSession Platform auth session. Treat `secret` as a credential — the
// trusted BFF stores it server-side (HTTP-only cookie), never in the browser.
type AuthSession struct {
    // The session id. Send it back as `session_id` to log out, or to have
    // `/auth/me` check that the session is still alive.
    Id string `json:"$id"`
    // When the session stops being valid on its own.
    Expire string `json:"expire"`
    // How the session was created. Server-minted sessions from this route are not
    // the browser-facing email/password ones, so this says which mechanism issued
    // it.
    Provider string `json:"provider"`
    // The session CREDENTIAL. Whoever holds it is logged in — the BFF keeps it
    // server-side (an HTTP-only cookie), never in the browser and never in a log.
    Secret string `json:"secret"`
    // The platform user this session belongs to — the `user_id` every other
    // auth route takes. NOT the contact id: the contact is in `contact`.
    UserId string `json:"userId"`

    // Used by Decode() method
    data []byte
}

func (model AuthSession) New(data []byte) *AuthSession {
    model.data = data
    return &model
}

func (model *AuthSession) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}