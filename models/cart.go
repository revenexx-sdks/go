package models

import (
    "encoding/json"
    "errors"
)

// Model
type Cart struct {
    // 
    AbandonedAt string `json:"abandoned_at"`
    // 
    ChannelId string `json:"channel_id"`
    // 
    ContactId string `json:"contact_id"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Currency string `json:"currency"`
    // 
    Id string `json:"id"`
    // 
    IsCurrent bool `json:"is_current"`
    // 
    ItemCount int `json:"item_count"`
    // 
    MarketId string `json:"market_id"`
    // 
    MergedIntoCartId string `json:"merged_into_cart_id"`
    // 
    Metadata interface{} `json:"metadata"`
    // 
    Name string `json:"name"`
    // 
    OrderRef string `json:"order_ref"`
    // 
    OrderedAt string `json:"ordered_at"`
    // 
    SessionKey string `json:"session_key"`
    // 
    Status string `json:"status"`
    // 
    Subtotal float64 `json:"subtotal"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model Cart) New(data []byte) *Cart {
    model.data = data
    return &model
}

func (model *Cart) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}