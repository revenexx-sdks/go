package models

import (
    "encoding/json"
    "errors"
)

// SegmentMemberCreateRequest Add one organization to a segment. Use
// source='manual' (the default) for hand-picked members; rule members are
// materialized by the recompute route.
type SegmentMemberCreateRequest struct {
    // The member company. Segments group companies, never people — a person is
    // reached through their organization.
    OrganizationId string `json:"organization_id"`
    // The segment.
    SegmentId string `json:"segment_id"`
    // How this membership came about: 'manual' is hand-picked, 'rule' was
    // materialized by a recompute. The distinction is load-bearing — a
    // recompute only ever inserts and deletes 'rule' rows, so a hand-picked
    // member survives every rule change. Default 'manual'.
    Source string `json:"source"`

    // Used by Decode() method
    data []byte
}

func (model SegmentMemberCreateRequest) New(data []byte) *SegmentMemberCreateRequest {
    model.data = data
    return &model
}

func (model *SegmentMemberCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}