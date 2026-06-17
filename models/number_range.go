package models

import (
    "encoding/json"
    "errors"
)

// Model
type NumberRange struct {
    // 
    ChannelId string `json:"channel_id"`
    // 
    Code string `json:"code"`
    // 
    Counter int `json:"counter"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    Metadata interface{} `json:"metadata"`
    // 
    Padding int `json:"padding"`
    // 
    PositionStep int `json:"position_step"`
    // 
    Prefix string `json:"prefix"`
    // 
    Step int `json:"step"`
    // 
    Suffix string `json:"suffix"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model NumberRange) New(data []byte) *NumberRange {
    model.data = data
    return &model
}

func (model *NumberRange) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}