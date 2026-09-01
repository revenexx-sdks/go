package models

import (
    "encoding/json"
    "errors"
)

// OrganizationMetricsRefreshResponse model.
type OrganizationMetricsRefreshResponse struct {
    // The instant the rolling windows are measured from. Send it back on every
    // continuation — that is what stops the 30/90/365-day windows sliding while
    // a multi-call refresh runs.
    AsOf string `json:"as_of"`
    // False if an insert had to fall back to row-at-a-time. A performance fact,
    // not an error.
    Batched bool `json:"batched"`
    // Rollup calls made to the orders app — the cross-app cost of this pass.
    Batches int `json:"batches"`
    // Where to resume: the id of the last organization this call processed. Send
    // it back verbatim; null when the pass finished. No example is published —
    // the value names a row in THIS tenant, and `cursor: "sample cursor"` reaches
    // PostgREST as a malformed uuid and comes back as a 400 nobody can read.
    Cursor string `json:"cursor"`
    // False means the budget ran out with work left — POST again with the
    // returned `cursor` AND `as_of`.
    Done bool `json:"done"`
    // Metrics rows created — organizations that had none yet.
    Inserted int `json:"inserted"`
    // Orders the orders app counted while answering this call.
    OrdersScanned int `json:"orders_scanned"`
    // Orders the orders app could not attribute to a company (B2C/guest). They
    // belong to no organization and land in no metrics row.
    OrdersWithoutOrganization int `json:"orders_without_organization"`
    // Organizations processed by THIS call.
    Organizations int `json:"organizations"`
    // Rows that already said the same thing — no write was issued. A routine
    // refresh is almost all of these.
    Unchanged int `json:"unchanged"`
    // Metrics rows whose numbers actually changed.
    Updated int `json:"updated"`
    // Of those, how many have at least one counted order.
    WithOrders int `json:"with_orders"`

    // Used by Decode() method
    data []byte
}

func (model OrganizationMetricsRefreshResponse) New(data []byte) *OrganizationMetricsRefreshResponse {
    model.data = data
    return &model
}

func (model *OrganizationMetricsRefreshResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}