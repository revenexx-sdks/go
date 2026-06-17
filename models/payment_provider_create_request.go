package models

import (
    "encoding/json"
    "errors"
)

// ActivatesACatalogPSPForThisTenantProvidersAreConfigurationNotCode Model
type PaymentProviderCreateRequest struct {
    // PSP credentials — the catalog's credential_fields say which keys the auth
    // scheme expects.
    Credentials interface{} `json:"credentials"`
    // Only enabled providers transact (default false).
    Enabled bool `json:"enabled"`
    // Display name — defaults to the catalog label.
    Name string `json:"name"`
    // Free-form provider options.
    Options interface{} `json:"options"`
    // Provider code — must exist in the catalog (GET
    // /payments/providers/catalog).
    Provider string `json:"provider"`
    // Sandbox/test credentials (default true).
    TestMode bool `json:"test_mode"`
    // Shared secret for PSP callback verification.
    WebhookSecret string `json:"webhook_secret"`

    // Used by Decode() method
    data []byte
}

func (model PaymentProviderCreateRequest) New(data []byte) *PaymentProviderCreateRequest {
    model.data = data
    return &model
}

func (model *PaymentProviderCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}