package models

import (
    "encoding/json"
    "errors"
)

// OrderCancelRequest Cancels the WHOLE order, and only while nothing has
// shipped. Both fields are optional unless the tenant requires a reason.
type OrderCancelRequest struct {
    // Who cancelled, as the caller reported it — an operator, a desk, a system.
    // Free text; this app does not resolve it against a user directory.
    CancelledBy string `json:"cancelled_by"`
    // Why it was cancelled, free text. Mandatory when the tenant sets
    // cancel_requires_reason — for those merchants an unexplained cancellation
    // is refused with a 400.
    Reason string `json:"reason"`

    // Used by Decode() method
    data []byte
}

func (model OrderCancelRequest) New(data []byte) *OrderCancelRequest {
    model.data = data
    return &model
}

func (model *OrderCancelRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}