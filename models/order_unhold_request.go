package models

import (
    "encoding/json"
    "errors"
)

// OrderUnholdRequest No payload — releasing the hold is a pure state
// transition, and it clears hold_reason with it. Send {}.
type OrderUnholdRequest struct {

    // Used by Decode() method
    data []byte
}

func (model OrderUnholdRequest) New(data []byte) *OrderUnholdRequest {
    model.data = data
    return &model
}

func (model *OrderUnholdRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}