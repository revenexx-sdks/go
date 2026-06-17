package models

import (
    "encoding/json"
    "errors"
)

// MessageList Model
type MessageList struct {
    // List of messages.
    Messages []Message `json:"messages"`
    // Total number of messages that matched your query.
    Total int `json:"total"`

    // Used by Decode() method
    data []byte
}

func (model MessageList) New(data []byte) *MessageList {
    model.data = data
    return &model
}

func (model *MessageList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}