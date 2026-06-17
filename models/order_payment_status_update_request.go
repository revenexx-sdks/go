package models

import (
    "encoding/json"
    "errors"
)

// Model
type OrderPaymentStatusUpdateRequest struct {
    // Reference into the payment system — merged into the order's payment
    // snapshot.
    PaymentId string `json:"payment_id"`
    // The new payment dimension value.
    Status string `json:"status"`

    // Used by Decode() method
    data []byte
}

func (model OrderPaymentStatusUpdateRequest) New(data []byte) *OrderPaymentStatusUpdateRequest {
    model.data = data
    return &model
}

func (model *OrderPaymentStatusUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}