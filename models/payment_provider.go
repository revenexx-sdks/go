package models

import (
    "encoding/json"
    "errors"
)

// PaymentProvider model.
type PaymentProvider struct {
    // When this PSP was configured for the tenant.
    CreatedAt string `json:"created_at"`
    // Only an enabled provider takes NEW payments: a method pointing at a
    // disabled one falls through to the tenant's `fallback_provider`, and to a
    // 422 if there is none. Nothing else reads it — capture, cancel and refund
    // on the payments this PSP already holds go on working — which is what
    // makes disabling the safe retirement and deleting the refused one.
    Enabled bool `json:"enabled"`
    // Id of the PSP configuration row — what the provider routes address. The
    // provider itself is named by `provider`.
    Id string `json:"id"`
    // Operator-facing name of the configuration. Defaults to the catalog label,
    // and is worth changing when a tenant runs two accounts with one PSP.
    Name string `json:"name"`
    // Per-provider switches this app understands, plus anything the merchant
    // keeps beside them. Three keys are the app's own: `logo_url` (the bundled
    // logo, filled in when the provider is seeded), `capture_method` and
    // `three_ds` (what the prism driver does today). Free jsonb — an unknown
    // key is stored and ignored.
    Options interface{} `json:"options"`
    // The catalog code of the PSP this row configures — one row per provider
    // per tenant. GET /payments/providers/catalog lists every code that may
    // appear here. It is what every payment and every method naming this PSP
    // resolves it by, so changing it is refused with 409 for as long as one of
    // them does.
    Provider string `json:"provider"`
    // Whether the driver talks to the PSP's sandbox. New configurations start in
    // test mode: a provider nobody verified must not touch live money.
    TestMode bool `json:"test_mode"`
    // When its configuration last changed — including a credential rotation,
    // which is otherwise invisible from the outside.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model PaymentProvider) New(data []byte) *PaymentProvider {
    model.data = data
    return &model
}

func (model *PaymentProvider) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}