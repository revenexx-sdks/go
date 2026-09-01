package models

import (
    "encoding/json"
    "errors"
)

// SegmentRuleCondition model.
type SegmentRuleCondition struct {
    // What the organization IS: an organizations column (name, status, vat_id,
    // branche, external_team_id) or 'setting:<key>' for a top-level key of
    // organizations.settings. Or what it DID, read from the organization_metrics
    // projection: order_count, order_count_30d/90d/365d, revenue_total,
    // revenue_30d/90d/365d, avg_order_value, avg_order_value_365d,
    // first_order_at, last_order_at, currency — plus the virtual
    // days_since_last_order (gt/gte/lt/lte only), which compares last_order_at
    // against a cut-off computed at evaluation time and never matches an
    // organization that never ordered (use last_order_at is_empty for those).
    Field string `json:"field"`
    // How `value` is compared to `field`. `contains`/`starts_with`/`ends_with`
    // are text matches; `in` takes an array; `is_empty`/`is_not_empty` take no
    // value at all.
    Operator string `json:"operator"`
    // Omitted for is_empty/is_not_empty; an array for 'in'; a string, number or
    // boolean otherwise. A number or boolean makes a 'setting:' condition compare
    // as JSONB, so it only matches values stored as a JSON number/boolean.
    Value string `json:"value"`

    // Used by Decode() method
    data []byte
}

func (model SegmentRuleCondition) New(data []byte) *SegmentRuleCondition {
    model.data = data
    return &model
}

func (model *SegmentRuleCondition) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}