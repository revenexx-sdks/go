package models

import (
    "encoding/json"
    "errors"
)

// Model
type PriceList struct {
    // 
    ChannelId string `json:"channel_id"`
    // 
    Code string `json:"code"`
    // 
    ContactId string `json:"contact_id"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Currency string `json:"currency"`
    // 
    Description string `json:"description"`
    // 
    Id string `json:"id"`
    // 
    IsDefault bool `json:"is_default"`
    // 
    Labels interface{} `json:"labels"`
    // 
    MarketId string `json:"market_id"`
    // 
    Metadata interface{} `json:"metadata"`
    // 
    Name string `json:"name"`
    // 
    OrganizationId string `json:"organization_id"`
    // 
    Priority int `json:"priority"`
    // 
    Status string `json:"status"`
    // 
    TaxIncluded bool `json:"tax_included"`
    // 
    UpdatedAt string `json:"updated_at"`
    // 
    ValidFrom string `json:"valid_from"`
    // 
    ValidUntil string `json:"valid_until"`

    // Used by Decode() method
    data []byte
}

func (model PriceList) New(data []byte) *PriceList {
    model.data = data
    return &model
}

func (model *PriceList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}