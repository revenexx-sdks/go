package models

import (
    "encoding/json"
    "errors"
)

// PaymentTransitionRequest model.
type PaymentTransitionRequest struct {
    // The operator's own words for why. Kept on the payment
    // (`metadata.cancel_reason` / `metadata.refund_reason`) AND handed to the
    // provider's own cancellation or refund reason field, so it is readable in
    // the PSP's dashboard too. Trimmed and cut at 500 characters.
    Reason string `json:"reason"`

    // Used by Decode() method
    data []byte
}

func (model PaymentTransitionRequest) New(data []byte) *PaymentTransitionRequest {
    model.data = data
    return &model
}

func (model *PaymentTransitionRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}