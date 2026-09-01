package models

import (
    "encoding/json"
    "errors"
)

// ValidationFailedResponse model.
type ValidationFailedResponse struct {
    // 
    Errors []string `json:"errors"`
    // 
    Status string `json:"status"`

    // Used by Decode() method
    data []byte
}

func (model ValidationFailedResponse) New(data []byte) *ValidationFailedResponse {
    model.data = data
    return &model
}

func (model *ValidationFailedResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}