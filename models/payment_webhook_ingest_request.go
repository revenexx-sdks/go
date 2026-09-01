package models

import (
    "encoding/json"
    "errors"
)

// PaymentWebhookIngestRequest The dispatch envelope from
// webhooks.revenexx.com. Nothing is required and nothing is constrained —
// three keys are read, and the rest is carried along.
type PaymentWebhookIngestRequest struct {
    // The dispatcher's delivery id. Echoed back as `delivery_id` so a delivery
    // and what the ledger did can be correlated.
    Id string `json:"id"`
    // The captured HTTP request as the PSP sent it.
    Request string `json:"request"`
    // Whether the ingress verified the callback signature against the provider's
    // `webhook_secret`. An explicit false is refused with 422: an endpoint may
    // run in annotate mode, and the ledger stays sovereign over one that does.
    Verified string `json:"verified"`

    // Used by Decode() method
    data []byte
}

func (model PaymentWebhookIngestRequest) New(data []byte) *PaymentWebhookIngestRequest {
    model.data = data
    return &model
}

func (model *PaymentWebhookIngestRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}