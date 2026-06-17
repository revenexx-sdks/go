package models

import (
    "encoding/json"
    "errors"
)

// Model
type CartItem struct {
    // 
    CartId string `json:"cart_id"`
    // 
    Configuration interface{} `json:"configuration"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Currency string `json:"currency"`
    // 
    Id string `json:"id"`
    // 
    LineTotal float64 `json:"line_total"`
    // 
    Metadata interface{} `json:"metadata"`
    // 
    Name string `json:"name"`
    // 
    Position int `json:"position"`
    // 
    ProductId string `json:"product_id"`
    // 
    Quantity float64 `json:"quantity"`
    // 
    Sku string `json:"sku"`
    // 
    Snapshot interface{} `json:"snapshot"`
    // 
    TaxRate float64 `json:"tax_rate"`
    // 
    Type string `json:"type"`
    // 
    Unit string `json:"unit"`
    // 
    UnitPrice float64 `json:"unit_price"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model CartItem) New(data []byte) *CartItem {
    model.data = data
    return &model
}

func (model *CartItem) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}