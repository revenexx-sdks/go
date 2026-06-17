package models

import (
    "encoding/json"
    "errors"
)

// APositionQuantityToCancelGuardedAgainstTheOpenUnshippedUncancelledQuantity
// Model
type OrderCancelPosition struct {
    // The order item (position) to act on.
    OrderItemId string `json:"order_item_id"`
    // Defaults to the full remaining quantity of the position.
    Quantity float64 `json:"quantity"`

    // Used by Decode() method
    data []byte
}

func (model OrderCancelPosition) New(data []byte) *OrderCancelPosition {
    model.data = data
    return &model
}

func (model *OrderCancelPosition) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}