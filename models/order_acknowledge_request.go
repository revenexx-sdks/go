package models

import (
    "encoding/json"
    "errors"
)

// OrderAcknowledgeRequest The acknowledgement carries one field, and it is
// optional: sending {} still stamps acknowledged_at, which is the point of
// the call. acknowledged_at is the server's clock and is never taken from the
// body.
type OrderAcknowledgeRequest struct {
    // The FULFILLING system's reference for this order, typically the ERP order
    // number. Written once by POST /orders/{id}/acknowledge and null until an
    // integration acknowledged it. Keeps the existing value when omitted.
    ExternalRef string `json:"external_ref"`

    // Used by Decode() method
    data []byte
}

func (model OrderAcknowledgeRequest) New(data []byte) *OrderAcknowledgeRequest {
    model.data = data
    return &model
}

func (model *OrderAcknowledgeRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}