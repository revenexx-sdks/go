package models

import (
    "encoding/json"
    "errors"
)

// OrderHoldRequest Stop the order. The reason is optional but is what the
// guard quotes back at whoever tries to ship, so an unexplained hold is a
// hold nobody can resolve.
type OrderHoldRequest struct {
    // Why the order is held, in the words the shipping guard quotes back. Null
    // when it is not held — releasing a hold clears it.
    Reason string `json:"reason"`

    // Used by Decode() method
    data []byte
}

func (model OrderHoldRequest) New(data []byte) *OrderHoldRequest {
    model.data = data
    return &model
}

func (model *OrderHoldRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}