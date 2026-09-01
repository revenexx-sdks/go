package models

import (
    "encoding/json"
    "errors"
)

// OrderCancellation A record of what was taken off an order and why —
// either the whole order (while nothing had shipped) or named quantities off
// a partly shipped one.
type OrderCancellation struct {
    // Who cancelled, as the caller reported it — an operator, a desk, a system.
    // Free text; this app does not resolve it against a user directory.
    CancelledBy string `json:"cancelled_by"`
    // When the cancellation was recorded.
    CreatedAt string `json:"created_at"`
    // Primary key of the cancellation record.
    Id string `json:"id"`
    // The order that was cancelled from.
    OrderId string `json:"order_id"`
    // What this record removed. A scope 'order' record carries every position in
    // full; a scope 'items' record carries exactly the quantities that were
    // named.
    Positions []OrderCancellationPosition `json:"positions"`
    // Why it was cancelled, free text. Mandatory when the tenant sets
    // cancel_requires_reason — for those merchants an unexplained cancellation
    // is refused with a 400.
    Reason string `json:"reason"`
    // Which of the two cancellations this was: 'order' is the full cancel (only
    // possible while nothing has shipped, and it cancels every position in full),
    // 'items' is the quantity-based one that takes open quantities off a partly
    // shipped order.
    Scope string `json:"scope"`

    // Used by Decode() method
    data []byte
}

func (model OrderCancellation) New(data []byte) *OrderCancellation {
    model.data = data
    return &model
}

func (model *OrderCancellation) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}