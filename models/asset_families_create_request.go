package models

import (
    "encoding/json"
    "errors"
)

// Model
type AssetFamiliesCreateRequest struct {
    // 
    Code string `json:"code"`
    // 
    Labels interface{} `json:"labels"`
    // 
    NamingConvention interface{} `json:"naming_convention"`

    // Used by Decode() method
    data []byte
}

func (model AssetFamiliesCreateRequest) New(data []byte) *AssetFamiliesCreateRequest {
    model.data = data
    return &model
}

func (model *AssetFamiliesCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}