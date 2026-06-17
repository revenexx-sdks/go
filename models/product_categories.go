package models

import (
    "encoding/json"
    "errors"
)

// Model
type ProductCategories struct {
    // 
    CategoryId string `json:"category_id"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    Position int `json:"position"`
    // 
    ProductId string `json:"product_id"`

    // Used by Decode() method
    data []byte
}

func (model ProductCategories) New(data []byte) *ProductCategories {
    model.data = data
    return &model
}

func (model *ProductCategories) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}