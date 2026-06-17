package models

import (
    "encoding/json"
    "errors"
)

// Model
type ProductsBatchRequest struct {
    // 
    Ids []string `json:"ids"`
    // 
    Skus []string `json:"skus"`

    // Used by Decode() method
    data []byte
}

func (model ProductsBatchRequest) New(data []byte) *ProductsBatchRequest {
    model.data = data
    return &model
}

func (model *ProductsBatchRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}