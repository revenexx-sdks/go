package models

import (
    "encoding/json"
    "errors"
)

// CartMergeResult Which cart survived, and what it cost. `target` is the cart
// that SURVIVES, already recomputed — that is the one to render. The source
// cart still exists and still holds its own lines: a merge copies them into
// the target and closes the source, it does not move them.
type CartMergeResult struct {
    // The source cart, now status merged, with merged_into_cart_id pointing at
    // the target. It still exists and still holds its own lines: the merge
    // copies, it does not move.
    MergedCartId string `json:"merged_cart_id"`
    // Lines read out of the source. Identical product lines at the same price add
    // up rather than duplicating, so the target may have gained fewer rows than
    // this.
    MergedLines int `json:"merged_lines"`
    // 
    Target Cart `json:"target"`

    // Used by Decode() method
    data []byte
}

func (model CartMergeResult) New(data []byte) *CartMergeResult {
    model.data = data
    return &model
}

func (model *CartMergeResult) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}