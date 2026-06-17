package models

import (
    "encoding/json"
    "errors"
)

// Model
type OrderReturnCompleteRequest struct {
    // How the return was settled (refund, replacement, …).
    Resolution string `json:"resolution"`

    // Used by Decode() method
    data []byte
}

func (model OrderReturnCompleteRequest) New(data []byte) *OrderReturnCompleteRequest {
    model.data = data
    return &model
}

func (model *OrderReturnCompleteRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}