package models

import (
    "encoding/json"
    "errors"
)

// Model
type OrderItem struct {
    // 
    Configuration interface{} `json:"configuration"`
    // 
    CostCenter string `json:"cost_center"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    LineTotal float64 `json:"line_total"`
    // 
    Metadata interface{} `json:"metadata"`
    // 
    Name string `json:"name"`
    // 
    OrderId string `json:"order_id"`
    // 
    Position int `json:"position"`
    // 
    PositionText string `json:"position_text"`
    // 
    Product interface{} `json:"product"`
    // 
    ProductId string `json:"product_id"`
    // 
    Quantity float64 `json:"quantity"`
    // 
    QuantityCancelled float64 `json:"quantity_cancelled"`
    // 
    QuantityReturned float64 `json:"quantity_returned"`
    // 
    QuantityShipped float64 `json:"quantity_shipped"`
    // 
    Sku string `json:"sku"`
    // 
    TaxAmount float64 `json:"tax_amount"`
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
    // 
    UserData interface{} `json:"user_data"`

    // Used by Decode() method
    data []byte
}

func (model OrderItem) New(data []byte) *OrderItem {
    model.data = data
    return &model
}

func (model *OrderItem) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}