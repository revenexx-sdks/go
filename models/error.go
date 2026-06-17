package models

import (
    "encoding/json"
    "errors"
)

// UniformGatewayErrorResponse Model
type Error struct {
    // 
    Error bool `json:"error"`
    // 
    Message string `json:"message"`

    // Used by Decode() method
    data []byte
}

func (model Error) New(data []byte) *Error {
    model.data = data
    return &model
}

func (model *Error) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}