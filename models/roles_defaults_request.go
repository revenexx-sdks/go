package models

import (
    "encoding/json"
    "errors"
)

// RolesDefaultsRequest No fields — send {}.
type RolesDefaultsRequest struct {

    // Used by Decode() method
    data []byte
}

func (model RolesDefaultsRequest) New(data []byte) *RolesDefaultsRequest {
    model.data = data
    return &model
}

func (model *RolesDefaultsRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}