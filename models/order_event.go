package models

import (
    "encoding/json"
    "errors"
)

// OrderEvent One entry of the audit trail, which is also the domain event
// feed: every action writes a row, the manifest emits order_event.created on
// insert, and the row name IS the event name on the bus.
type OrderEvent struct {
    // Who caused it: the resolved contact id of the acting principal. Only
    // order.placed and order.requested carry one today — every other row is
    // null — so filtering on it filters to those two names. The database
    // constrains nothing here (the column is text); the uuid shape is what this
    // app WRITES, which is also why no example is published: no id an app invents
    // names a row a tenant holds.
    Actor string `json:"actor"`
    // When it happened. The trail comes back oldest first, which is the order a
    // human reads a history in.
    CreatedAt string `json:"created_at"`
    // Primary key of the event row.
    Id string `json:"id"`
    // WHAT happened, and this is the domain event: the manifest emits
    // order_event.created on insert and this value is the event name on the bus.
    // The names this app writes are order.placed, order.requested, order.updated,
    // order.acknowledged, order.cancelled, order.item.cancelled,
    // order.shipment.created, order.completed, order.held, order.unheld,
    // order.payment_status.changed, order.comment.added, order.return.registered,
    // order.return.received, order.return.completed and order.return.rejected.
    Name string `json:"name"`
    // The order this happened to.
    OrderId string `json:"order_id"`
    // The machine-readable body, and its shape follows `name`. order.placed /
    // order.requested carry number, grand_total, currency, item_count, cart_id
    // — plus approval_reason (permission | value_threshold) and threshold when
    // the order is waiting for sign-off. order.shipment.created carries
    // shipment_id, number, carrier, tracking_code and the booked positions.
    // order.item.cancelled and order.return.* carry positions and the reason or
    // resolution. order.completed carries via (shipment | payment | manual).
    // order.payment_status.changed carries from, to and payment_id. Nothing
    // validates it: it is what the route that wrote the row put there.
    Payload interface{} `json:"payload"`

    // Used by Decode() method
    data []byte
}

func (model OrderEvent) New(data []byte) *OrderEvent {
    model.data = data
    return &model
}

func (model *OrderEvent) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}