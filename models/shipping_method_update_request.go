package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type ShippingMethodUpdateRequest struct {
    // Carrier anchor for the upcoming carrier connect (dynamic rates, tracking
    // links).
    Carrier string `json:"carrier"`
    // Stable method code, unique per tenant (e.g. standard, express).
    Code string `json:"code"`
    // Allowed ISO 3166-1 alpha-2 codes; null or empty = worldwide.
    Countries []string `json:"countries"`
    // ISO 4217 code (default EUR).
    Currency string `json:"currency"`
    // 
    Description string `json:"description"`
    // Only enabled methods appear in rate responses (default false).
    Enabled bool `json:"enabled"`
    // Delivery-time estimate for the checkout (days, upper bound).
    EtaDaysMax int `json:"eta_days_max"`
    // Delivery-time estimate for the checkout (days, lower bound).
    EtaDaysMin int `json:"eta_days_min"`
    // Free shipping at or above this order value — wins over every pricing
    // model.
    FreeAbove float64 `json:"free_above"`
    // Localized display names keyed by locale (e.g. {de, en}).
    Labels interface{} `json:"labels"`
    // Attribute name for matrix_basis 'attribute'.
    MatrixAttribute string `json:"matrix_attribute"`
    // The measure a matrix method prices over; 'attribute' reads matrix_attribute
    // from the rate request.
    MatrixBasis string `json:"matrix_basis"`
    // Free-form metadata.
    Metadata interface{} `json:"metadata"`
    // Display name.
    Name string `json:"name"`
    // Sort order in the checkout (default 0).
    Position int `json:"position"`
    // The fixed price (default 0) — ignored for 'free' and 'matrix'.
    Price float64 `json:"price"`
    // Pricing model (default 'fixed'): one price, no price, or tiered over a
    // measure.
    PricingType string `json:"pricing_type"`

    // Used by Decode() method
    data []byte
}

func (model ShippingMethodUpdateRequest) New(data []byte) *ShippingMethodUpdateRequest {
    model.data = data
    return &model
}

func (model *ShippingMethodUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}