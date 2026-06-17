package models

import (
    "encoding/json"
    "errors"
)

// Metric Model
type Metric struct {
    // The date at which this metric was aggregated in ISO 8601 format.
    Date string `json:"date"`
    // The value of this metric at the timestamp.
    Value int `json:"value"`

    // Used by Decode() method
    data []byte
}

func (model Metric) New(data []byte) *Metric {
    model.data = data
    return &model
}

func (model *Metric) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}