package models

import (
    "encoding/json"
    "errors"
)

// NoPayloadReceivingIsAPureStateTransitionRegisteredReceived Model
type OrderReturnReceiveRequest struct {

    // Used by Decode() method
    data []byte
}

func (model OrderReturnReceiveRequest) New(data []byte) *OrderReturnReceiveRequest {
    model.data = data
    return &model
}

func (model *OrderReturnReceiveRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}