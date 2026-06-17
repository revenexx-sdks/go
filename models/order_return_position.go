package models

import (
    "encoding/json"
    "errors"
)

// APositionQuantityToReturnGuardedAgainstTheShippedNotYetReturnedQuantity
// Model
type OrderReturnPosition struct {
    // The order item (position) to act on.
    OrderItemId string `json:"order_item_id"`
    // Defaults to the full remaining quantity of the position.
    Quantity float64 `json:"quantity"`
    // Report this position for restocking when the return completes (the explicit
    // inventories.restock call stays with the orchestrator).
    Restock bool `json:"restock"`

    // Used by Decode() method
    data []byte
}

func (model OrderReturnPosition) New(data []byte) *OrderReturnPosition {
    model.data = data
    return &model
}

func (model *OrderReturnPosition) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}