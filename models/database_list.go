package models

import (
    "encoding/json"
    "errors"
)

// DatabasesList Model
type DatabaseList struct {
    // List of databases.
    Databases []Database `json:"databases"`
    // Total number of databases that matched your query.
    Total int `json:"total"`

    // Used by Decode() method
    data []byte
}

func (model DatabaseList) New(data []byte) *DatabaseList {
    model.data = data
    return &model
}

func (model *DatabaseList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}