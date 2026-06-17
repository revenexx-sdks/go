package models

import (
    "encoding/json"
    "errors"
)

// Model
type StockLevel struct {
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    LocationId string `json:"location_id"`
    // 
    Metadata interface{} `json:"metadata"`
    // 
    OnHand float64 `json:"on_hand"`
    // 
    ProductId string `json:"product_id"`
    // 
    ReorderPoint float64 `json:"reorder_point"`
    // 
    Reserved float64 `json:"reserved"`
    // 
    Sku string `json:"sku"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model StockLevel) New(data []byte) *StockLevel {
    model.data = data
    return &model
}

func (model *StockLevel) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}