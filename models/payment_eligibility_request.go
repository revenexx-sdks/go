package models

import (
    "encoding/json"
    "errors"
)

// TheBuyerContextRestrictionDimensionsAreANDedEntriesWithinADimensionORedEmptyUnrestricted
// Model
type PaymentEligibilityRequest struct {
    // Order amount the fees are computed against (default 0).
    Amount float64 `json:"amount"`
    // Buyer ISO country code — methods with country restrictions need it.
    Country string `json:"country"`
    // ISO 4217 code (default EUR).
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