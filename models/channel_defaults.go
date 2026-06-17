package models

import (
    "encoding/json"
    "errors"
)

// Model
type ChannelDefaults struct {
    // Channel codes created by this call.
    Created []string `json:"created"`
    // Default channel codes that already existed.
    Existing []string `json:"existing"`

    // Used by Decode() method
    data []byte
}

func (model ChannelDefaults) New(data []byte) *ChannelDefaults {
    model.data = data
    return &model
}

func (model *ChannelDefaults) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}