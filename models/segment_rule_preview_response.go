package models

import (
    "encoding/json"
    "errors"
)

// SegmentRulePreviewResponse model.
type SegmentRulePreviewResponse struct {
    // The cap that applied (5000), or null when the rule was answered by a single
    // count query and no cap was needed.
    Cap int `json:"cap"`
    // True when the combined evaluation hit the id cap, which makes `count` a
    // lower bound.
    Capped bool `json:"capped"`
    // How many organizations the rule selects. Exact when 'capped' is false; a
    // LOWER BOUND when it is true.
    Count int `json:"count"`
    // How the conditions were combined for this preview.
    RuleMatch string `json:"rule_match"`
    // A handful of the organizations the rule selects — enough for an operator
    // to recognise whether the rule means what they thought. Never the full set.
    Sample []interface{} `json:"sample"`
    // The segment named in the path. It is not read — the rule comes from the
    // body — but it has to exist.
    SegmentId string `json:"segment_id"`
    // What the rule selects. Only 'organizations' exists.
    Target string `json:"target"`

    // Used by Decode() method
    data []byte
}

func (model SegmentRulePreviewResponse) New(data []byte) *SegmentRulePreviewResponse {
    model.data = data
    return &model
}

func (model *SegmentRulePreviewResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}