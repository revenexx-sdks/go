package models

import (
    "encoding/json"
    "errors"
)

// OrderPaymentStatusUpdateRequest model.
type OrderPaymentStatusUpdateRequest struct {
    // The reference into the payment system. MERGED into the order's payment
    // snapshot under 'payment_id' — the rest of the snapshot is left alone —
    // and carried in the order.payment_status.changed event. Omitted leaves the
    // snapshot untouched.
    PaymentId string `json:"payment_id"`
    // The new value of the payment dimension. Whether the order is PAID, and the
    // dimension this app does not decide: it is fed from outside through POST
    // /orders/{id}/payment-status (the payments app or an ERP), and only seeded
    // at place-time from payment.status. Orthogonal to the lifecycle — a
    // completed order can still be open, and a paid one can still be pending.
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