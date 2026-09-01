package models

import (
    "encoding/json"
    "errors"
)

// OrganizationMetrics What an organization has BOUGHT, materialized from the
// orders app. One row per organization — including all-zero rows for
// companies that never ordered, so a 'never bought anything' rule has
// something to match.
type OrganizationMetrics struct {
    // revenue_total / order_count, computed here from the sums rather than
    // averaged upstream. Zero when there are no orders.
    AvgOrderValue float64 `json:"avg_order_value"`
    // revenue_365d / order_count_365d. Zero when there were none in the window.
    AvgOrderValue365d float64 `json:"avg_order_value_365d"`
    // When this row was last written. The projection is materialized, so this is
    // how stale the numbers are.
    ComputedAt string `json:"computed_at"`
    // When the projection row first appeared.
    CreatedAt string `json:"created_at"`
    // The single ISO 4217 currency all counted orders were in. NULL when there
    // were none, and also when there were several — read `currency_mixed` to
    // tell those two apart.
    Currency string `json:"currency"`
    // True when this company ordered in more than one currency. The sums are
    // still stored (dropping money is worse), but they are not comparable against
    // a threshold, and a rule reading revenue should say so.
    CurrencyMixed bool `json:"currency_mixed"`
    // When this company first ordered. Null if it never has — that is what
    // makes it usable as "is this a customer at all?".
    FirstOrderAt string `json:"first_order_at"`
    // Primary key of the projection row.
    Id string `json:"id"`
    // When this company last ordered. Null if it never has, which is why the
    // virtual `days_since_last_order` rule field never matches those companies:
    // use `last_order_at is_empty` for them.
    LastOrderAt string `json:"last_order_at"`
    // Orders ever counted for this company.
    OrderCount int `json:"order_count"`
    // Orders in the 30 days before `orders_as_of`. A rolling window, not a
    // calendar month.
    OrderCount30d int `json:"order_count_30d"`
    // Orders in the 365 days before `orders_as_of`.
    OrderCount365d int `json:"order_count_365d"`
    // Orders in the 90 days before `orders_as_of`.
    OrderCount90d int `json:"order_count_90d"`
    // The instant the rolling windows were measured from. Pinned across a chunked
    // refresh, so a multi-call pass cannot let the windows slide underneath it.
    OrdersAsOf string `json:"orders_as_of"`
    // The company these numbers describe. One row per organization, and rows
    // exist for companies that never ordered — all zeros rather than missing,
    // so a "never bought" rule matches something.
    OrganizationId string `json:"organization_id"`
    // Revenue in the 30 days before `orders_as_of`.
    Revenue30d float64 `json:"revenue_30d"`
    // Revenue in the 365 days before `orders_as_of`. The usual "how big is this
    // customer" number, and the one a key-account rule should read.
    Revenue365d float64 `json:"revenue_365d"`
    // Revenue in the 90 days before `orders_as_of`.
    Revenue90d float64 `json:"revenue_90d"`
    // Revenue ever counted, in `currency`. Which orders count is the orders app's
    // decision, not this app's.
    RevenueTotal float64 `json:"revenue_total"`
    // The tenant this row belongs to — the store slug, not an id. Set by the
    // platform from the authenticated context, never by a caller; a write that
    // carries it is ignored, and no request can read another tenant's rows by
    // sending a different one.
    TenantId string `json:"tenant_id"`
    // When the row last changed. Unchanged numbers are not rewritten, so this can
    // lag `computed_at`.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model OrganizationMetrics) New(data []byte) *OrganizationMetrics {
    model.data = data
    return &model
}

func (model *OrganizationMetrics) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}