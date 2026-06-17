package models

import (
    "encoding/json"
    "errors"
)

// Model
type OrderReturnRejectRequest struct {
    // Fallback for 'resolution'.
    Reason string `json:"reason"`
    // Why the return was rejected.
    Resolution string `json:"resolution"`

    // Used by Decode() method
    data []byte
}

func (model OrderReturnRejectRequest) New(data []byte) *OrderReturnRejectRequest {
    model.data = data
    return &model
}

func (model *OrderReturnRejectRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}