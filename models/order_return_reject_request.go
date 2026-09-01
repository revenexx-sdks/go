package models

import (
    "encoding/json"
    "errors"
)

// OrderReturnRejectRequest model.
type OrderReturnRejectRequest struct {
    // Free-text fallback for 'resolution' — a sentence about this one return,
    // not a value out of the set.
    Reason string `json:"reason"`
    // Why the return was refused.
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