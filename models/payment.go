package models

import (
    "encoding/json"
    "errors"
)

// Model
type Payment struct {
    // 
    Amount float64 `json:"amount"`
    // 
    AuthorizedAt string `json:"authorized_at"`
    // 
    CapturedAt string `json:"captured_at"`
    // 
    CartId string `json:"cart_id"`
    // 
    ContactId string `json:"contact_id"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Currency string `json:"currency"`
    // 
    ErrorMessage string `json:"error_message"`
    // 
    FailedAt string `json:"failed_at"`
    // 
    FeeAmount float64 `json:"fee_amount"`
    // 
    Id string `json:"id"`
    // 
    IdempotencyKey string `json:"idempotency_key"`
    // 
    Kind string `json:"kind"`
    // 
    Metadata interface{} `json:"metadata"`
    // 
    MethodCode string `json:"method_code"`
    // 
    NextAction interface{} `json:"next_action"`
    // 
    OrderRef string `json:"order_ref"`
    // 
    Provider string `json:"provider"`
    // 
    PspPaymentId string `json:"psp_payment_id"`
    // 
    RefundedAt string `json:"refunded_at"`
    // 
    Status string `json:"status"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model Payment) New(data []byte) *Payment {
    model.data = data
    return &model
}

func (model *Payment) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}