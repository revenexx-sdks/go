package models

import (
    "encoding/json"
    "errors"
)

// OrderCompleteRequest No required fields — send {}.
type OrderCompleteRequest struct {
    // Who closed the order, as the caller reports it. Not stored on the order: it
    // is carried in the order.completed event's payload, which is where the audit
    // trail keeps who did what. Free text, not resolved against a user directory.
    CompletedBy string `json:"completed_by"`

    // Used by Decode() method
    data []byte
}

func (model OrderCompleteRequest) New(data []byte) *OrderCompleteRequest {
    model.data = data
    return &model
}

func (model *OrderCompleteRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}