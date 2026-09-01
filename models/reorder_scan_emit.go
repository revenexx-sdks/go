package models

import (
    "encoding/json"
    "errors"
)

// ReorderScanEmit model.
type ReorderScanEmit struct {
    // The event id on the bus. Stable per (row, day), which is what makes a
    // re-run harmless.
    EventId string `json:"event_id"`
    // The stock row the event is about.
    StockLevelId string `json:"stock_level_id"`

    // Used by Decode() method
    data []byte
}

func (model ReorderScanEmit) New(data []byte) *ReorderScanEmit {
    model.data = data
    return &model
}

func (model *ReorderScanEmit) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}