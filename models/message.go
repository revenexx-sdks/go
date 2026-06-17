package models

import (
    "encoding/json"
    "errors"
)

// Message Model
type Message struct {
    // Message creation time in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Message ID.
    Id string `json:"$id"`
    // Message update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Data of the message.
    Data interface{} `json:"data"`
    // The time when the message was delivered.
    DeliveredAt string `json:"deliveredAt"`
    // Number of recipients the message was delivered to.
    DeliveredTotal int `json:"deliveredTotal"`
    // Delivery errors if any.
    DeliveryErrors []string `json:"deliveryErrors"`
    // Message provider type.
    ProviderType string `json:"providerType"`
    // The scheduled time for message.
    ScheduledAt string `json:"scheduledAt"`
    // Status of delivery.
    Status string `json:"status"`
    // Target IDs set as recipients.
    Targets []string `json:"targets"`
    // Topic IDs set as recipients.
    Topics []string `json:"topics"`
    // User IDs set as recipients.
    Users []string `json:"users"`

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