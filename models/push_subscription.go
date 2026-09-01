package models

import (
    "encoding/json"
    "errors"
)

// PushSubscription model.
type PushSubscription struct {
    // 
    CreatedAt string `json:"created_at"`
    // 
    Endpoint string `json:"endpoint"`
    // 
    Id string `json:"id"`
    // 
    LastSeenAt string `json:"last_seen_at"`
    // 
    SubscriberId string `json:"subscriber_id"`
    // 
    TenantId string `json:"tenant_id"`
    // 
    UpdatedAt string `json:"updated_at"`
    // 
    UserAgent string `json:"user_agent"`

    // Used by Decode() method
    data []byte
}

func (model PushSubscription) New(data []byte) *PushSubscription {
    model.data = data
    return &model
}

func (model *PushSubscription) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}