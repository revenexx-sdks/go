package models

import (
    "encoding/json"
    "errors"
)

// OrderListKindMakeDefaultRequest No payload — send {}. The kind is named
// by the path, and there is nothing else to decide.
type OrderListKindMakeDefaultRequest struct {

    // Used by Decode() method
    data []byte
}

func (model OrderListKindMakeDefaultRequest) New(data []byte) *OrderListKindMakeDefaultRequest {
    model.data = data
    return &model
}

func (model *OrderListKindMakeDefaultRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}