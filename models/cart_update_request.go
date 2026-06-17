package models

import (
    "encoding/json"
    "errors"
)

// OnlySafeColumnsAreUpdatableStatusMovesThroughTheLifecycleRoutes Model
type CartUpdateRequest struct {
    // 
    ChannelId string `json:"channel_id"`
    // ISO 4217 code.
    Currency string `json:"currency"`
    // 
    MarketId string `json:"market_id"`
    // Free-form metadata.
    Metadata interface{} `json:"metadata"`
    // 
    Name string `json:"name"`

    // Used by Decode() method
    data []byte
}

func (model CartUpdateRequest) New(data []byte) *CartUpdateRequest {
    model.data = data
    return &model
}

func (model *CartUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}