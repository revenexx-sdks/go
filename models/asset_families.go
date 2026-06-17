package models

import (
    "encoding/json"
    "errors"
)

// Model
type AssetFamilies struct {
    // 
    Code string `json:"code"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    Labels interface{} `json:"labels"`
    // 
    NamingConvention interface{} `json:"naming_convention"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model AssetFamilies) New(data []byte) *AssetFamilies {
    model.data = data
    return &model
}

func (model *AssetFamilies) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}