package models

import (
    "encoding/json"
    "errors"
)

// TheDispatchEnvelopeFromWebhooksRevenexxComRequestBodyCarriesTheRawVendorShapedPSPCallbackStripePaymentIntentsOrTheGenericEventPspPaymentIdOrderRefErrorShapeIntentionallyUnconstrainedSoNoPSPNotificationIsEverRejectedAtTheGate
// Model
type PaymentWebhookIngestRequest struct {

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