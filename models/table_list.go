package models

import (
    "encoding/json"
    "errors"
)

// TablesList Model
type TableList struct {
    // List of tables.
    Tables []Table `json:"tables"`
    // Total number of tables that matched your query.
    Total int `json:"total"`

    // Used by Decode() method
    data []byte
}

func (model TableList) New(data []byte) *TableList {
    model.data = data
    return &model
}

func (model *TableList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}