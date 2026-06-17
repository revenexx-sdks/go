package models

import (
    "encoding/json"
    "errors"
)

// Model
type PriceEntry struct {
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    Metadata interface{} `json:"metadata"`
    // 
    PriceListId string `json:"price_list_id"`
    // 
    PriceType string `json:"price_type"`
    // 
    ProductId string `json:"product_id"`
    // 
    QuantityMin float64 `json:"quantity_min"`
    // 
    Sku string `json:"sku"`
    // 
    Unit string `json:"unit"`
    // 
    UnitPrice float64 `json:"unit_price"`
    // 
    UpdatedAt string `json:"updated_at"`
    // 
    ValidFrom string `json:"valid_from"`
    // 
    ValidUntil string `json:"valid_until"`

    // Used by Decode() method
    data []byte
}

func (model PriceEntry) New(data []byte) *PriceEntry {
    model.data = data
    return &model
}

func (model *PriceEntry) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}