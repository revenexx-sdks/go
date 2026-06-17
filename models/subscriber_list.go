package models

import (
    "encoding/json"
    "errors"
)

// SubscriberList Model
type SubscriberList struct {
    // List of subscribers.
    Subscribers []Subscriber `json:"subscribers"`
    // Total number of subscribers that matched your query.
    Total int `json:"total"`

    // Used by Decode() method
    data []byte
}

func (model SubscriberList) New(data []byte) *SubscriberList {
    model.data = data
    return &model
}

func (model *SubscriberList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}