package models

import (
    "encoding/json"
    "errors"
)

// ReorderScanRequest No fields — send `{}`. What counts as low follows each
// row's own `reorder_point` and the market's `reorder_point_default`, exactly
// as GET /inventories/reorder-alerts computes it.
type ReorderScanRequest struct {

    // Used by Decode() method
    data []byte
}

func (model ReorderScanRequest) New(data []byte) *ReorderScanRequest {
    model.data = data
    return &model
}

func (model *ReorderScanRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}