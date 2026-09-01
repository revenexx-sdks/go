package models

import (
    "encoding/json"
    "errors"
)

// CartMergeRequest model.
type CartMergeRequest struct {
    // The cart being folded in. It must be active, and it does NOT survive as a
    // workspace: its lines are copied into the target, it becomes status merged,
    // and merged_into_cart_id points at the target. Its own lines stay on it as
    // the record of what was moved.
    SourceCartId string `json:"source_cart_id"`
    // The cart that SURVIVES. Must be active; it gains the source's lines
    // (identical product lines at the same price adding up) and its totals are
    // recomputed.
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