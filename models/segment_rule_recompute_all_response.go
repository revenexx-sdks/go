package models

import (
    "encoding/json"
    "errors"
)

// SegmentRuleRecomputeAllResponse model.
type SegmentRuleRecomputeAllResponse struct {
    // Rule memberships inserted across every segment in THIS call.
    Added int `json:"added"`
    // False when any segment is unfinished or skipped — call again.
    Done bool `json:"done"`
    // Segments whose own recompute raised — they carry `error` and `status` in
    // `segments` and did not abort the run.
    Failed int `json:"failed"`
    // Ruled segments the run looked at.
    Processed int `json:"processed"`
    // Rule memberships deleted across every segment in THIS call.
    Removed int `json:"removed"`
    // One entry per segment; a failed segment carries `error` and `status`
    // instead of the counters.
    Segments []interface{} `json:"segments"`
    // Segments the budget did not reach at all.
    Skipped int `json:"skipped"`

    // Used by Decode() method
    data []byte
}

func (model SegmentRuleRecomputeAllResponse) New(data []byte) *SegmentRuleRecomputeAllResponse {
    model.data = data
    return &model
}

func (model *SegmentRuleRecomputeAllResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}