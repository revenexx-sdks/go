package models

import (
    "encoding/json"
    "errors"
)

// EligiblePaymentMethod One method as a checkout should render it: identity,
// wording, and what it costs this buyer.
type EligiblePaymentMethod struct {
    // The code to send back as `method_code` when the payment is created.
    Code string `json:"code"`
    // The currency `fee` is in — the one the request asked with, echoed.
    Currency string `json:"currency"`
    // The merchant's line about this method, to show beside it at checkout.
    Description string `json:"description"`
    // The surcharge this method costs THIS buyer, already computed against the
    // requested amount — a fixed fee as it stands, a percentage resolved into
    // an amount. Not a column: no CHECK bounds it, so none is declared.
    Fee float64 `json:"fee"`
    // How `fee` was arrived at, for a checkout that wants to show "2 % surcharge"
    // rather than the amount.
    FeeType string `json:"fee_type"`
    // Whether choosing this method starts a PSP flow ('psp') or authorizes
    // immediately ('self_managed').
    Kind string `json:"kind"`
    // Buyer-facing names keyed by language tag, or null when the merchant
    // configured none — then `name` is all there is.
    Labels interface{} `json:"labels"`
    // The operator-facing name. Prefer `labels` for anything a buyer reads.
    Name string `json:"name"`
    // The merchant's sort order. The list is already sorted by it; it is carried
    // so a client that re-sorts can put it back.
    Position int `json:"position"`
    // The PSP behind it, for a checkout that has to load a provider SDK before it
    // can collect an instrument. null for self-managed methods.
    Provider string `json:"provider"`

    // Used by Decode() method
    data []byte
}

func (model EligiblePaymentMethod) New(data []byte) *EligiblePaymentMethod {
    model.data = data
    return &model
}

func (model *EligiblePaymentMethod) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}