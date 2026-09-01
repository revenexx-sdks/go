package models

import (
    "encoding/json"
    "errors"
)

// SegmentRuleRecomputeRequest model.
type SegmentRuleRecomputeRequest struct {
    // Continuation token from a previous response — the id of the last
    // organization the pass touched. Omit to resume or start automatically; pass
    // null to force a restart from the beginning.
    Cursor string `json:"cursor"`

    // Used by Decode() method
    data []byte
}

func (model SegmentRuleRecomputeRequest) New(data []byte) *SegmentRuleRecomputeRequest {
    model.data = data
    return &model
}

func (model *SegmentRuleRecomputeRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}