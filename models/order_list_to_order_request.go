package models

import (
    "encoding/json"
    "errors"
)

// OrderListToOrderRequest Every field is optional — the buyer, the
// organization and the positions all come from the list.
type OrderListToOrderRequest struct {
    // ISO 4217 code. Omit to let the orders app apply the market default.
    Currency string `json:"currency"`
    // The BUYER's own order or purchase-order number, forwarded to the orders app
    // verbatim. Free text and never generated here: it exists so the paperwork
    // can carry the number the buyer's accounts payable will look for.
    CustomerOrderNumber string `json:"customer_order_number"`

    // Used by Decode() method
    data []byte
}

func (model OrderListToOrderRequest) New(data []byte) *OrderListToOrderRequest {
    model.data = data
    return &model
}

func (model *OrderListToOrderRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}