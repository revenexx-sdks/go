package models

import (
    "encoding/json"
    "errors"
)

// UnauthenticatedResponse model.
type UnauthenticatedResponse struct {
    // 
    Message string `json:"message"`

    // Used by Decode() method
    data []byte
}

func (model UnauthenticatedResponse) New(data []byte) *UnauthenticatedResponse {
    model.data = data
    return &model
}

func (model *UnauthenticatedResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}