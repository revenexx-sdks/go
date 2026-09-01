package models

import (
    "encoding/json"
    "errors"
)

// CartMergeIntoRequest model.
type CartMergeIntoRequest struct {
    // Receiving cart (must be active). The cart in the path is the source and
    // becomes status merged.
    TargetCartId string `json:"target_cart_id"`

    // Used by Decode() method
    data []byte
}

func (model CartMergeIntoRequest) New(data []byte) *CartMergeIntoRequest {
    model.data = data
    return &model
}

func (model *CartMergeIntoRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}