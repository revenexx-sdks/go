package models

import (
    "encoding/json"
    "errors"
)

// Message model.
type Message struct {
    // 
    Attachments []interface{} `json:"attachments"`
    // 
    Attempts int `json:"attempts"`
    // 
    BindingId string `json:"binding_id"`
    // 
    Channel string `json:"channel"`
    // 
    ClickCount int `json:"click_count"`
    // 
    ClickedAt string `json:"clicked_at"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Data []interface{} `json:"data"`
    // 
    DeliveredAt string `json:"delivered_at"`
    // 
    Error string `json:"error"`
    // 
    FromDraft bool `json:"from_draft"`
    // 
    Id string `json:"id"`
    // 
    IdempotencyFingerprint string `json:"idempotency_fingerprint"`
    // 
    IdempotencyKey string `json:"idempotency_key"`
    // 
    Locale string `json:"locale"`
    // 
    Market string `json:"market"`
    // 
    MessageClass string `json:"message_class"`
    // 
    OpenCount int `json:"open_count"`
    // 
    OpenedAt string `json:"opened_at"`
    // 
    ProviderMessageId string `json:"provider_message_id"`
    // 
    ScheduledFor string `json:"scheduled_for"`
    // 
    SentAt string `json:"sent_at"`
    // 
    SourceEventId string `json:"source_event_id"`
    // 
    Status string `json:"status"`
    // 
    Subject string `json:"subject"`
    // 
    SuppressionReason string `json:"suppression_reason"`
    // 
    TemplateKey string `json:"template_key"`
    // 
    TenantId string `json:"tenant_id"`
    // 
    To string `json:"to"`

    // Used by Decode() method
    data []byte
}

func (model Message) New(data []byte) *Message {
    model.data = data
    return &model
}

func (model *Message) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}