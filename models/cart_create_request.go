package models

import (
    "encoding/json"
    "errors"
)

// ACartNeedsAnOwnerContactIdCustomerOrSessionKeyGuest Model
type CartCreateRequest struct {
    // 
    ChannelId string `json:"channel_id"`
    // Owning customer contact.
    ContactId string `json:"contact_id"`
    // ISO 4217 code (default EUR).
    Currency string `json:"currency"`
    // Make this THE current cart of its owner.
    IsCurrent bool `json:"is_current"`
    // 
    MarketId string `json:"market_id"`
    // Free-form metadata.
    Metadata interface{} `json:"metadata"`
    // Display name (default 'Cart').
    Name string `json:"name"`
    // Owning guest session.
    SessionKey string `json:"session_key"`

    // Used by Decode() method
    data []byte
}

func (model CartCreateRequest) New(data []byte) *CartCreateRequest {
    model.data = data
    return &model
}

func (model *CartCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}