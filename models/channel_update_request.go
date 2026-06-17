package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type ChannelUpdateRequest struct {
    // Stable channel code, unique per tenant (e.g. shop, punchout-acme).
    Code string `json:"code"`
    // Mark as the default channel (default false).
    IsDefault bool `json:"is_default"`
    // Localized display names keyed by locale.
    Labels interface{} `json:"labels"`
    // Display name.
    Name string `json:"name"`
    // Sort position (default 0).
    Position int `json:"position"`
    // Lifecycle status (default 'active').
    Status string `json:"status"`
    // Where business happens (default 'storefront').
    Type string `json:"type"`

    // Used by Decode() method
    data []byte
}

func (model ChannelUpdateRequest) New(data []byte) *ChannelUpdateRequest {
    model.data = data
    return &model
}

func (model *ChannelUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}