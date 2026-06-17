package models

import (
    "encoding/json"
    "errors"
)

// CreatesANDAuthorizesSelfManagedMethodsAuthorizeImmediatelyPSPMethodsMayAnswerNextActionRedirectEligibilityIsReCheckedServerSide
// Model
type PaymentCreateRequest struct {
    // Order amount — 0 is legal (free orders), negative is not.
    Amount float64 `json:"amount"`
    // The cart this payment pays for.
    CartId string `json:"cart_id"`
    // Paying customer contact.
    ContactId string `json:"contact_id"`
    // Buyer ISO country code for the eligibility check.
    Country string `json:"country"`
    // ISO 4217 code (default EUR).
    Currency string `json:"currency"`
    // Same key answers the same payment instead of a duplicate.
    IdempotencyKey string `json:"idempotency_key"`
    // Free-form metadata.
    Metadata interface{} `json:"metadata"`
    // Code of a configured payment method.
    MethodCode string `json:"method_code"`
    // External order reference — also the webhook fallback key.
    OrderRef string `json:"order_ref"`
    // Where the PSP redirect flow returns the buyer to.
    ReturnUrl string `json:"return_url"`

    // Used by Decode() method
    data []byte
}

func (model PaymentCreateRequest) New(data []byte) *PaymentCreateRequest {
    model.data = data
    return &model
}

func (model *PaymentCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}