package models

import (
    "encoding/json"
    "errors"
)

// ShippingRatesRequest The buyer context the checkout resolves rates for —
// matrix methods need their measure (weight, quantity, order value or
// attribute) to apply.
type ShippingRatesRequest struct {
    // The instant to evaluate the delivery estimate at (ISO 8601). Omitted: now.
    // Lets a storefront compute the cut-off in its own timezone.
    At string `json:"at"`
    // Measure values for attribute matrices, keyed by attribute NAME — the key
    // a matrix method names in its matrix_attribute, and the value the number its
    // tiers are matched against. Summed over the basket by the caller, not by
    // this app. Only the key a method asks for is read; anything else in the map
    // is carried along and ignored, and a value that is not a finite number
    // excludes that method with a reason rather than failing the quote.
    Attributes interface{} `json:"attributes"`
    // Destination ISO 3166-1 alpha-2 code — compared upper-cased against method
    // and carrier country restrictions. Omitted or null: every method that
    // restricts by country is excluded, with a reason.
    Country string `json:"country"`
    // ISO 4217 code, echoed into the rates (default 'EUR'). Echoed, not
    // converted: this app prices in the currency the method carries.
    Currency string `json:"currency"`
    // Buyer market for tax resolution. Omitted: the market matching `country`,
    // else the tenant's sole market — never an arbitrary one.
    MarketId string `json:"market_id"`
    // Order value (default 0) — drives order_value matrices, and free-above
    // thresholds when no sided value is sent. Read on the basis the tenant's
    // free_above_compares setting declares.
    OrderValue float64 `json:"order_value"`
    // Order value including tax. Compared against free-above thresholds when
    // free_above_compares is 'gross'.
    OrderValueGross float64 `json:"order_value_gross"`
    // Order value excluding tax. Compared against free-above thresholds when
    // free_above_compares is 'net'.
    OrderValueNet float64 `json:"order_value_net"`
    // Total quantity — measure for quantity matrices.
    Quantity float64 `json:"quantity"`
    // Total weight — measure for weight matrices. Read in weight_unit and
    // converted to the unit the tiers are keyed in.
    Weight float64 `json:"weight"`
    // The unit `weight` is expressed in, as a CODE into the tenant's own weight
    // units (GET /shipping/weight-units). Omitted, it is the unit this market
    // quotes in. A unit the tenant does not keep is a 400 — a mis-read weight
    // prices the wrong bracket silently, and guessing is worse than refusing.
    WeightUnit string `json:"weight_unit"`

    // Used by Decode() method
    data []byte
}

func (model ShippingRatesRequest) New(data []byte) *ShippingRatesRequest {
    model.data = data
    return &model
}

func (model *ShippingRatesRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}