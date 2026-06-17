package models

import (
    "encoding/json"
    "errors"
)

// ExecutionsList Model
type ExecutionList struct {
    // List of executions.
    Executions []Execution `json:"executions"`
    // Total number of executions that matched your query.
    Total int `json:"total"`

    // Used by Decode() method
    data []byte
}

func (model ExecutionList) New(data []byte) *ExecutionList {
    model.data = data
    return &model
}

func (model *ExecutionList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}