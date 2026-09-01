package models

import (
    "encoding/json"
    "errors"
)

// SegmentCreateRequest model.
type SegmentCreateRequest struct {
    // Stable identifier, unique per tenant — what other apps and integrations
    // name the segment by. Free text, but lowercase with underscores is the
    // convention every seeded vocabulary follows.
    Code string `json:"code"`
    // Localized display names keyed by language tag. Null means nobody translated
    // it and a client falls back to showing the code.
    Labels interface{} `json:"labels"`
    // Sort order in the cockpit, ascending. Ties fall back to insertion order.
    // Default 0.
    Position int `json:"position"`
    // How the conditions combine: 'all' (default) is AND, 'any' is OR. Null means
    // the same as 'all'.
    RuleMatch string `json:"rule_match"`
    // The selector that decides membership, stored verbatim. Null means the
    // segment is manual-only. The same rule language product categories use,
    // evaluated over organization columns, `setting:<key>` entries and the
    // organization_metrics projection — so 'no order in 365 days' is
    // expressible without joining the orders app. Null makes the segment
    // manual-only. Changing it does not move a single membership — run the
    // recompute.
    Rules SegmentRules `json:"rules"`

    // Used by Decode() method
    data []byte
}

func (model SegmentCreateRequest) New(data []byte) *SegmentCreateRequest {
    model.data = data
    return &model
}

func (model *SegmentCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}