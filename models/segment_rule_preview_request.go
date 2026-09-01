package models

import (
    "encoding/json"
    "errors"
)

// SegmentRulePreviewRequest model.
type SegmentRulePreviewRequest struct {
    // The conditions, combined by `rule_match`. At least one, at most 25.
    Conditions []SegmentRuleCondition `json:"conditions"`
    // How the conditions combine. Default 'all'.
    RuleMatch string `json:"rule_match"`
    // Only 'organizations' is supported; any other value is rejected. A segment
    // groups COMPANIES — the people are reached through them.
    Target string `json:"target"`

    // Used by Decode() method
    data []byte
}

func (model SegmentRulePreviewRequest) New(data []byte) *SegmentRulePreviewRequest {
    model.data = data
    return &model
}

func (model *SegmentRulePreviewRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}