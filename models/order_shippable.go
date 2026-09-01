package models

import (
    "encoding/json"
    "errors"
)

// OrderShippable What a shipment of this order may still contain, and whether
// one would be accepted at all — answered by the same code POST
// /orders/{id}/ship runs, so the two cannot drift.
type OrderShippable struct {
    // Why not, in the very words POST /orders/{id}/ship would refuse with —
    // including the hold reason where there is one. Null when `shippable` is
    // true.
    BlockedReason string `json:"blocked_reason"`
    // How many positions still have an open quantity — the number of lines a
    // shipment dialog would offer.
    OpenPositions int `json:"open_positions"`
    // The summed open quantity over those positions. Mixes units where the order
    // does, so it is a headline figure, not a total to act on.
    OpenQuantity float64 `json:"open_quantity"`
    // Just enough of the order to render the answer — the full row is GET
    // /orders/{id}.
    Order OrderShippableOrder `json:"order"`
    // Every position of the order, in position order, each with its open
    // quantity.
    Positions []OrderShippablePosition `json:"positions"`
    // Whether a shipment would be accepted RIGHT NOW — the one question a
    // "create shipment" button should be enabled on. False when the order is
    // held, cancelled, completed, or has nothing open.
    Shippable bool `json:"shippable"`

    // Used by Decode() method
    data []byte
}

func (model OrderShippable) New(data []byte) *OrderShippable {
    model.data = data
    return &model
}

func (model *OrderShippable) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}