package models

import (
    "encoding/json"
    "errors"
)

// RuntimesList Model
type RuntimeList struct {
    // List of runtimes.
    Runtimes []Runtime `json:"runtimes"`
    // Total number of runtimes that matched your query.
    Total int `json:"total"`

    // Used by Decode() method
    data []byte
}

func (model RuntimeList) New(data []byte) *RuntimeList {
    model.data = data
    return &model
}

func (model *RuntimeList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}