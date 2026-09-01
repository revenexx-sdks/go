package models

import (
    "encoding/json"
    "errors"
)

// OrderItemsCancelRequest model.
type OrderItemsCancelRequest struct {
    // Who cancelled, as the caller reported it — an operator, a desk, a system.
    // Free text; this app does not resolve it against a user directory.
    CancelledBy string `json:"cancelled_by"`
    // The quantities to take off the order. Required here, unlike on /ship and
    // /return: cancelling everything by default is not a thing anybody should be
    // able to do by omission — that is what /cancel is for.
    Positions []OrderCancelPosition `json:"positions"`
    // Why it was cancelled, free text. Mandatory when the tenant sets
    // cancel_requires_reason — for those merchants an unexplained cancellation
    // is refused with a 400.
    Reason string `json:"reason"`

    // Used by Decode() method
    data []byte
}

func (model OrderItemsCancelRequest) New(data []byte) *OrderItemsCancelRequest {
    model.data = data
    return &model
}

func (model *OrderItemsCancelRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}