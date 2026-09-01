package models

import (
    "encoding/json"
    "errors"
)

// CartConversionReservation What this app ASKED inventories for, and what it
// answered. This app holds no stock: inventories picks the location, applies
// the backorder policy and owns the hold's expiry.
type CartConversionReservation struct {
    // Lines inventories accepted without stock behind them, under the tenant's
    // backorder policy — its policy, not this app's.
    Backordered int `json:"backordered"`
    // inventories' hold deadline — its TTL, not this app's.
    ExpiresAt string `json:"expires_at"`
    // A hold exists. False with `requested: true` means inventories was asked and
    // refused — `reason` says why, and only convert_reserves_stock = require
    // turns that into a 409.
    Ok bool `json:"ok"`
    // The reference the reservation was booked under: the `order_ref` of the
    // request, or the cart id when the call carried none. This is the string to
    // hand inventories when releasing the hold.
    OrderRef string `json:"order_ref"`
    // Why no hold exists — stated, never implied. Present whenever `ok` is
    // false, and also on the never case.
    Reason string `json:"reason"`
    // False when convert_reserves_stock is 'never' — no call was made at all,
    // which is reported rather than dressed up as a silent success.
    Requested bool `json:"requested"`
    // Lines inventories confirmed a hold for.
    Reservations int `json:"reservations"`
    // The HTTP status inventories answered with, present only when it refused.
    // 404 is its own case: the tenant has no inventories app at all, which is a
    // different problem from not enough stock.
    Status int `json:"status"`

    // Used by Decode() method
    data []byte
}

func (model CartConversionReservation) New(data []byte) *CartConversionReservation {
    model.data = data
    return &model
}

func (model *CartConversionReservation) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}