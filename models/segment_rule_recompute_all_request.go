package models

import (
    "encoding/json"
    "errors"
)

// SegmentRuleRecomputeAllRequest No parameters — send {}.
type SegmentRuleRecomputeAllRequest struct {

    // Used by Decode() method
    data []byte
}

func (model SegmentRuleRecomputeAllRequest) New(data []byte) *SegmentRuleRecomputeAllRequest {
    model.data = data
    return &model
}

func (model *SegmentRuleRecomputeAllRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}