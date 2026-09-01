package models

import (
    "encoding/json"
    "errors"
)

// OrganizationMetricsFreshness model.
type OrganizationMetricsFreshness struct {
    // Companies with no metrics row yet. A rule reading revenue silently skips
    // them, so this is the number to watch after an import.
    Missing int `json:"missing"`
    // The OLDEST computed_at in the table — the floor, not an average. Null
    // when there are no rows at all.
    OldestComputedAt string `json:"oldest_computed_at"`
    // The anchor those oldest numbers were measured from.
    OrdersAsOf string `json:"orders_as_of"`
    // Companies in this tenant.
    Organizations int `json:"organizations"`
    // Metrics rows that exist — at most one per company.
    Rows int `json:"rows"`

    // Used by Decode() method
    data []byte
}

func (model OrganizationMetricsFreshness) New(data []byte) *OrganizationMetricsFreshness {
    model.data = data
    return &model
}

func (model *OrganizationMetricsFreshness) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}