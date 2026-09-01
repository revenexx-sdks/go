package models

import (
    "encoding/json"
    "errors"
)

// ReorderAlerts model.
type ReorderAlerts struct {
    // The rows at or below their reorder point, worst first (by `shortfall`).
    // Computed on read, so it is never stale — and never empty because of
    // caching: an empty list means nothing is low, unless `enabled` is false.
    Alerts []ReorderAlert `json:"alerts"`
    // false when reorder_alert_enabled is off — the list is then empty by
    // policy, not because nothing is low.
    Enabled bool `json:"enabled"`
    // The threshold applied to rows carrying none of their own.
    ReorderPointDefault float64 `json:"reorder_point_default"`

    // Used by Decode() method
    data []byte
}

func (model ReorderAlerts) New(data []byte) *ReorderAlerts {
    model.data = data
    return &model
}

func (model *ReorderAlerts) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}