package models

import (
    "encoding/json"
    "errors"
)

// CustomersDefaultsRequest No fields — send {}.
type CustomersDefaultsRequest struct {

    // Used by Decode() method
    data []byte
}

func (model CustomersDefaultsRequest) New(data []byte) *CustomersDefaultsRequest {
    model.data = data
    return &model
}

func (model *CustomersDefaultsRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}