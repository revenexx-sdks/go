package models

import (
    "encoding/json"
    "errors"
)

// TheBuyerContextTheCheckoutResolvesRatesForMatrixMethodsNeedTheirMeasureWeightQuantityOrderValueOrAttributeToApply
// Model
type ShippingRatesRequest struct {
    // Measure values for attribute matrices, keyed by attribute name.
    Attributes interface{} `json:"attributes"`
    // Destination ISO 3166-1 alpha-2 code — checked against method country
    // restrictions.
    Country string `json:"country"`
    // Echoed into the rates (default 'EUR').
    Currency string `json:"currency"`
    // Buyer market for tax resolution (else inferred from country, else first
    // market).
    MarketId string `json:"market_id"`
    // Order value (default 0) — drives free-above thresholds and order_value
    // matrices.
    OrderValue float64 `json:"order_value"`
    // Total quantity — measure for quantity matrices.
    Quantity float64 `json:"quantity"`
    // Total weight — measure for weight matrices.
    Weight float64 `json:"weight"`

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