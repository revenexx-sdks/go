package models

import (
    "encoding/json"
    "errors"
)

// ShippingRatesBasis How this answer was measured — the tenant settings
// that shaped it, echoed so the numbers can be re-derived.
type ShippingRatesBasis struct {
    // The instant the delivery estimates were computed from.
    EvaluatedAt string `json:"evaluated_at"`
    // Whether free-above thresholds were compared against the net or the gross
    // order value.
    FreeAboveCompares string `json:"free_above_compares"`
    // The measure a matrix method without its own basis priced over.
    MatrixBasisDefault string `json:"matrix_basis_default"`
    // The unit the request expressed its weight in; converted to weight_unit
    // before any tier was matched.
    RequestWeightUnit string `json:"request_weight_unit"`
    // Kilograms per unit of `request_weight_unit`, as applied.
    RequestWeightUnitFactor float64 `json:"request_weight_unit_factor"`
    // The unit the rate tiers are keyed in — this market's `weight_unit`
    // setting, else the unit the tenant flagged as default.
    WeightUnit string `json:"weight_unit"`
    // Kilograms per unit of `weight_unit`, as applied. Echoed because a unit is a
    // code PLUS a number and the number is what priced the parcel — a quote has
    // to be re-derivable from its own payload, not from a table the merchant may
    // since have edited.
    WeightUnitFactor float64 `json:"weight_unit_factor"`

    // Used by Decode() method
    data []byte
}

func (model ShippingRatesBasis) New(data []byte) *ShippingRatesBasis {
    model.data = data
    return &model
}

func (model *ShippingRatesBasis) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}