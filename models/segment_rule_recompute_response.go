package models

import (
    "encoding/json"
    "errors"
)

// SegmentRuleRecomputeResponse model.
type SegmentRuleRecomputeResponse struct {
    // Rule memberships inserted by THIS call.
    Added int `json:"added"`
    // True when every membership insert used a bulk array request; false if any
    // batch fell back to row-at-a-time.
    Batched bool `json:"batched"`
    // Set when the pass completes.
    ComputedAt string `json:"computed_at"`
    // Send back on the next call; null when the pass is done.
    Cursor string `json:"cursor"`
    // False means work remains — POST again with `cursor`.
    Done bool `json:"done"`
    // Matching organizations examined by THIS call.
    Processed int `json:"processed"`
    // Rule memberships deleted by THIS call.
    Removed int `json:"removed"`
    // The segment that was recomputed.
    SegmentId string `json:"segment_id"`
    // The rule's full match count; null until done.
    Total int `json:"total"`

    // Used by Decode() method
    data []byte
}

func (model SegmentRuleRecomputeResponse) New(data []byte) *SegmentRuleRecomputeResponse {
    model.data = data
    return &model
}

func (model *SegmentRuleRecomputeResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}