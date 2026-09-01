package models

import (
    "encoding/json"
    "errors"
)

// SegmentMember One organization inside one segment, and the record of how it
// got there (hand-picked or matched by the rule).
type SegmentMember struct {
    // When the organization joined the segment.
    CreatedAt string `json:"created_at"`
    // Primary key of the membership row.
    Id string `json:"id"`
    // The member company. Segments group companies, never people — a person is
    // reached through their organization.
    OrganizationId string `json:"organization_id"`
    // The segment.
    SegmentId string `json:"segment_id"`
    // How this membership came about: 'manual' is hand-picked, 'rule' was
    // materialized by a recompute. The distinction is load-bearing — a
    // recompute only ever inserts and deletes 'rule' rows, so a hand-picked
    // member survives every rule change.
    Source string `json:"source"`
    // The tenant this row belongs to — the store slug, not an id. Set by the
    // platform from the authenticated context, never by a caller; a write that
    // carries it is ignored, and no request can read another tenant's rows by
    // sending a different one.
    TenantId string `json:"tenant_id"`

    // Used by Decode() method
    data []byte
}

func (model SegmentMember) New(data []byte) *SegmentMember {
    model.data = data
    return &model
}

func (model *SegmentMember) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}