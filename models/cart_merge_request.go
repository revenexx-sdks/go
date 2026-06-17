package models

import (
    "encoding/json"
    "errors"
)

// Model
type CartMergeRequest struct {
    // Cart whose lines move into the target (becomes status merged).
    SourceCartId string `json:"source_cart_id"`
    // Receiving cart (must be active).
    TargetCartId string `json:"target_cart_id"`

    // Used by Decode() method
    data []byte
}

func (model CartMergeRequest) New(data []byte) *CartMergeRequest {
    model.data = data
    return &model
}

func (model *CartMergeRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}