package models

import (
    "encoding/json"
    "errors"
)

// OrderShippableOrder Just enough of the order to render the answer — the
// full row is GET /orders/{id}.
type OrderShippableOrder struct {
    // Whether the order has SHIPPED, and the one dimension nobody writes: it is
    // DERIVED after every quantity change from the positions' own bookkeeping.
    // 'fulfilled' means shipped >= ordered − cancelled across all positions,
    // 'partial' means something went out. Sending it has no effect; ship, cancel
    // or return something and it moves.
    FulfillmentStatus string `json:"fulfillment_status"`
    // Why the order is held, in the words the shipping guard quotes back. Null
    // when it is not held — releasing a hold clears it.
    HoldReason string `json:"hold_reason"`
    // The order this answer is about.
    Id string `json:"id"`
    // The order number a human quotes — drawn from the tenant's order range at
    // place-time, unique per tenant and never reused. It is NOT the id: every
    // route addresses an order by uuid, and GET /orders?number=… is how a
    // number becomes one.
    Number string `json:"number"`
    // A business stop, ORTHOGONAL to status: a held order keeps its lifecycle
    // state and is refused at the guards. How far the hold reaches is the
    // tenant's call (on_hold_blocks: shipping only, shipping and cancellation, or
    // nothing at all).
    OnHold bool `json:"on_hold"`
    // Where the order stands in its LIFECYCLE, and one of three independent
    // status dimensions. 'pending' = created but not placed, an order waiting for
    // approval; 'placed' = accepted, nothing shipped; 'in_fulfillment' = part of
    // it has gone out, or all of it has and the tenant does not close on
    // shipment; 'completed' and 'cancelled' end it. Moved by the action routes
    // only — it is not writable through PUT /orders/{id}.
    Status string `json:"status"`

    // Used by Decode() method
    data []byte
}

func (model OrderShippableOrder) New(data []byte) *OrderShippableOrder {
    model.data = data
    return &model
}

func (model *OrderShippableOrder) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}