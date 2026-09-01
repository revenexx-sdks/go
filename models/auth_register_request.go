package models

import (
    "encoding/json"
    "errors"
)

// AuthRegisterRequest model.
type AuthRegisterRequest struct {
    // The buyer's address. It becomes the login AND the unique key of the
    // contact, so a second registration with it is a 409 — including while the
    // first one is still waiting for approval.
    Email string `json:"email"`
    // Given name. Optional: an ERP import often has only a mailbox.
    FirstName string `json:"first_name"`
    // Family name. Optional for the same reason.
    LastName string `json:"last_name"`
    // The language this person is written to in — BCP 47, and one of the
    // store's configured locales. Null falls back to the store default. One of
    // the store's own locales, or the call is a 400.
    Locale string `json:"locale"`
    // JOIN an existing company — the invite shape. Neither
    // b2b_registration_enabled nor b2c_registration_enabled applies to it.
    OrganizationId string `json:"organization_id"`
    // FOUND a new company, with this contact as its admin. This is what makes the
    // registration a B2B one; leaving it out registers a standalone buyer.
    OrganizationName string `json:"organization_name"`
    // The password the buyer chooses. It is hashed by the identity service at
    // this moment and never travels again: an approval later enables the account,
    // it does not issue a new credential.
    Password string `json:"password"`
    // Where the welcome mail's button points — the buyer's first stop in this
    // shop. Absent, the mail still goes out and simply carries no button. Ignored
    // when the registration is an APPLICATION: there is no account to send
    // anybody to yet.
    Url string `json:"url"`
    // VAT identification number (USt-IdNr. in Germany) — the closest thing a
    // B2B buyer has to a legal identity. Validated against the EU VIES service
    // when the tenant's `organization_vat_id_required` setting is on, and stored
    // verbatim otherwise, including for buyers outside the EU. Required when the
    // tenant's `organization_vat_id_required` is on, and checked BEFORE the
    // company is created so a bad one leaves no half-founded organization behind.
    VatId string `json:"vat_id"`
    // Where the address-confirmation link points, when the tenant's
    // `email_verification` asks for one on registration. `userId`, `secret` and
    // `expire` are appended, and `PUT /customers/auth/verification` takes the
    // first two. Without it the registration still succeeds and
    // `verification_sent` is false — this app cannot invent a storefront URL,
    // and a link pointing nowhere is worse than none.
    VerificationUrl string `json:"verification_url"`

    // Used by Decode() method
    data []byte
}

func (model AuthRegisterRequest) New(data []byte) *AuthRegisterRequest {
    model.data = data
    return &model
}

func (model *AuthRegisterRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}