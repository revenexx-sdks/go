package models

import (
    "encoding/json"
    "errors"
)

// AMethodNeedsItsIdentityCodeName Model
type PaymentMethodCreateRequest struct {
    // Stable method code (unique per tenant, e.g. 'invoice', 'card').
    Code string `json:"code"`
    // Allowed ISO country codes — empty/omitted = unrestricted.
    Countries []string `json:"countries"`
    // 
    Description string `json:"description"`
    // Disabled methods are never eligible (default false).
    Enabled bool `json:"enabled"`
    // Fixed amount or percent value, per fee_type (default 0).
    FeeAmount float64 `json:"fee_amount"`
    // ISO 4217 code (default EUR).
    FeeCurrency string `json:"fee_currency"`
    // How 'fee_amount' applies (default 'none').
    FeeType string `json:"fee_type"`
    // Self-managed (merchant fulfils, default) or PSP-backed ('provider' required
    // to transact).
    Kind string `json:"kind"`
    // Localized display names ({ de, en, … }).
    Labels interface{} `json:"labels"`
    // Maximum order amount — omitted = no upper bound.
    MaxOrderValue float64 `json:"max_order_value"`
    // Free-form metadata.
    Metadata interface{} `json:"metadata"`
    // Minimum order amount — omitted = no lower bound.
    MinOrderValue float64 `json:"min_order_value"`
    // Display name.
    Name string `json:"name"`
    // Sort position in the checkout (default 0).
    Position int `json:"position"`
    // PSP code from the catalog — only for kind 'psp'.
    Provider string `json:"provider"`
    // The provider's payment method id (e.g. 'card', 'paypal').
    ProviderMethod string `json:"provider_method"`

    // Used by Decode() method
    data []byte
}

func (model PaymentMethodCreateRequest) New(data []byte) *PaymentMethodCreateRequest {
    model.data = data
    return &model
}

func (model *PaymentMethodCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}