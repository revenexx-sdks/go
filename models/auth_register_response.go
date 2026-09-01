package models

import (
    "encoding/json"
    "errors"
)

// AuthRegisterResponse model.
type AuthRegisterResponse struct {
    // True when the tenant runs registration_mode='approval_required' — do NOT
    // log the buyer in.
    ApprovalRequired bool `json:"approval_required"`
    // The stored customer record — this app is its system of record.
    Contact Contact `json:"contact"`
    // 'pending' means the login is disabled until a merchant approves.
    RegistrationStatus string `json:"registration_status"`
    // The platform user that was created. Keep it: logout, /auth/me and the
    // recovery confirm all take it.
    UserId string `json:"user_id"`
    // Whether an address confirmation went out. True only when the tenant's
    // `email_verification` asks for one on registration, the registration is a
    // finished account rather than an application, and `verification_url` was
    // supplied.
    VerificationSent bool `json:"verification_sent"`
    // Whether the tenant's welcome mail went out. Best effort on purpose: the
    // account exists either way, and a registration is not undone because a
    // message service was unreachable. False for an APPLICATION, which is not an
    // account yet and is announced by `registration.submitted` instead.
    WelcomeSent bool `json:"welcome_sent"`

    // Used by Decode() method
    data []byte
}

func (model AuthRegisterResponse) New(data []byte) *AuthRegisterResponse {
    model.data = data
    return &model
}

func (model *AuthRegisterResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}