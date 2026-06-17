package models

import (
    "encoding/json"
    "errors"
)

// Model
type StockMovement struct {
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    LocationId string `json:"location_id"`
    // 
    Metadata interface{} `json:"metadata"`
    // 
    OrderRef string `json:"order_ref"`
    // 
    ProductId string `json:"product_id"`
    // 
    Quantity float64 `json:"quantity"`
    // 
    Reason string `json:"reason"`
    // 
    Sku string `json:"sku"`
    // 
    Type string `json:"type"`

    // Used by Decode() method
    data []byte
}

func (model StockMovement) New(data []byte) *StockMovement {
    model.data = data
    return &model
}

func (model *StockMovement) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}