package models

import (
    "encoding/json"
    "errors"
)

// OrderCustomerRollup Additive order facts for one organization. Average
// order value is revenue_total / order_count.
type OrderCustomerRollup struct {
    // Every currency seen on the counted orders, sorted. MORE THAN ONE MEANS THE
    // SUMS MIX CURRENCIES — nothing here converts, so a two-currency row's
    // revenue_total is a sum of unlike numbers and should be shown per currency
    // or not at all.
    Currencies []string `json:"currencies"`
    // When this company first ordered — placed_at where there is one, otherwise
    // created_at. Null cannot happen on a row that exists, but the field is
    // nullable because the columns behind it are.
    FirstOrderAt string `json:"first_order_at"`
    // When they last ordered. Together with as_of this is the recency a churn
    // rule reads.
    LastOrderAt string `json:"last_order_at"`
    // How many orders of this company were counted — orders in one of the
    // counted statuses, over all time.
    OrderCount int `json:"order_count"`
    // Orders in the 30 days before as_of.
    OrderCount30d int `json:"order_count_30d"`
    // Orders in the 365 days before as_of — the rolling year a "still active"
    // rule usually asks about.
    OrderCount365d int `json:"order_count_365d"`
    // Orders in the 90 days before as_of.
    OrderCount90d int `json:"order_count_90d"`
    // The company these facts belong to — the id the customers app knows it by.
    // Every row of the answer carries one; orders without an organization are
    // counted in orders_without_organization instead.
    OrganizationId string `json:"organization_id"`
    // Revenue in the 30 days before as_of.
    Revenue30d float64 `json:"revenue_30d"`
    // Revenue in the 365 days before as_of.
    Revenue365d float64 `json:"revenue_365d"`
    // Revenue in the 90 days before as_of.
    Revenue90d float64 `json:"revenue_90d"`
    // Sum of grand_total over the counted orders. Gross: it includes tax and
    // shipping, because grand_total does.
    RevenueTotal float64 `json:"revenue_total"`

    // Used by Decode() method
    data []byte
}

func (model OrderCustomerRollup) New(data []byte) *OrderCustomerRollup {
    model.data = data
    return &model
}

func (model *OrderCustomerRollup) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}