package models

import (
    "encoding/json"
    "errors"
)

// PaymentEligibilityRequest The buyer context — restriction dimensions are
// ANDed, entries within a dimension ORed, empty = unrestricted.
type PaymentEligibilityRequest struct {
    // The order amount the order-value bounds are checked against and the
    // percentage fees are computed from. Defaults to 0, which excludes every
    // method carrying a minimum. Nothing is written, so the ledger's own amount
    // bound does not apply here.
    Amount float64 `json:"amount"`
    // The buyer's ISO 3166-1 alpha-2 country code. A method restricted to
    // countries is excluded without it — an unknown buyer sees only the
    // unrestricted methods, which is the safe default and not a bug.
    Country string `json:"country"`
    // ISO 4217 code the amount is in, echoed onto every computed fee. Defaults to
    // EUR. This app does no conversion: the fee comes back in the currency it was
    // asked with.
    Currency string `json:"currency"`

    // Used by Decode() method
    data []byte
}

func (model PaymentEligibilityRequest) New(data []byte) *PaymentEligibilityRequest {
    model.data = data
    return &model
}

func (model *PaymentEligibilityRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}