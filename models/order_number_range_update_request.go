package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type OrderNumberRangeUpdateRequest struct {
    // 
    ChannelId string `json:"channel_id"`
    // Range key drawn by the app ('order', 'delivery', 'return') — unique per
    // tenant.
    Code string `json:"code"`
    // Current counter value (default 0) — the next number draws counter+step.
    Counter int `json:"counter"`
    // Free-form metadata.
    Metadata interface{} `json:"metadata"`
    // Zero-padding width of the counter (default 6).
    Padding int `json:"padding"`
    // Position numbering increment for order items (default 10).
    PositionStep int `json:"position_step"`
    // Default ''.
    Prefix string `json:"prefix"`
    // Counter increment per drawn number (default 1).
    Step int `json:"step"`
    // Default ''.
    Suffix string `json:"suffix"`

    // Used by Decode() method
    data []byte
}

func (model OrderNumberRangeUpdateRequest) New(data []byte) *OrderNumberRangeUpdateRequest {
    model.data = data
    return &model
}

func (model *OrderNumberRangeUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}