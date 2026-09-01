package models

import (
    "encoding/json"
    "errors"
)

// ReorderScan model.
type ReorderScan struct {
    // One entry per published event, in the order they went out. Re-running the
    // scan on the same day returns the SAME ids and publishes nothing a second
    // time — the event id is derived from the row and the day, and the bus
    // drops the repeat.
    Emitted []ReorderScanEmit `json:"emitted"`
    // false when reorder_alert_enabled is off — nothing was published, and not
    // because nothing is low.
    Enabled bool `json:"enabled"`
    // How many rows were at or below their point when the scan ran.
    Scanned int `json:"scanned"`

    // Used by Decode() method
    data []byte
}

func (model ReorderScan) New(data []byte) *ReorderScan {
    model.data = data
    return &model
}

func (model *ReorderScan) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}