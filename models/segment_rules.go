package models

import (
    "encoding/json"
    "errors"
)

// SegmentRules The selector that decides membership, stored verbatim. Null
// means the segment is manual-only. The same rule language product categories
// use, evaluated over organization columns, `setting:<key>` entries and the
// organization_metrics projection — so 'no order in 365 days' is
// expressible without joining the orders app. Null makes the segment
// manual-only. Changing it does not move a single membership — run the
// recompute.
type SegmentRules struct {
    // The conditions, combined by `rule_match`. At least one, at most 25.
    Conditions []SegmentRuleCondition `json:"conditions"`
    // Only 'organizations' is supported; any other value is rejected. A segment
    // groups COMPANIES — the people are reached through them.
    Target string `json:"target"`

    // Used by Decode() method
    data []byte
}

func (model SegmentRules) New(data []byte) *SegmentRules {
    model.data = data
    return &model
}

func (model *SegmentRules) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}