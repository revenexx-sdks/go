package models

import (
    "encoding/json"
    "errors"
)

// NoPayloadReleasingTheHoldIsAPureStateTransition Model
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