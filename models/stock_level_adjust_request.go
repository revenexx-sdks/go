package models

import (
    "encoding/json"
    "errors"
)

// StockLevelAdjustRequest Correct ONE stock row. The row already knows its
// location and its item, so a caller owes only the signed delta and a reason
// — which is exactly what an operator can be asked for in a dialog.
type StockLevelAdjustRequest struct {
    // The SIGNED correction to this row's `on_hand`: −3 writes off three, +3
    // finds three. A delta, not the new balance. Zero is refused (400). A
    // correction that would take `on_hand` below zero is a 422 the database
    // insists on; one that would take it below this row's own `reserved` is a 422
    // the `allow_negative_stock` setting can permit.
    Quantity float64 `json:"quantity"`
    // Why this row is being corrected, written onto the ledger booking. Owed
    // unless `movement_reason_required` is 'none'.
    Reason string `json:"reason"`

    // Used by Decode() method
    data []byte
}

func (model StockLevelAdjustRequest) New(data []byte) *StockLevelAdjustRequest {
    model.data = data
    return &model
}

func (model *StockLevelAdjustRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}