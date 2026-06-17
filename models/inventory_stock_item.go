package models

import (
    "encoding/json"
    "errors"
)

// AnItemAndItsQuantityProductIdOrSku Model
type InventoryStockItem struct {
    // 
    ProductId string `json:"product_id"`
    // 
    Quantity float64 `json:"quantity"`
    // 
    Sku string `json:"sku"`

    // Used by Decode() method
    data []byte
}

func (model InventoryStockItem) New(data []byte) *InventoryStockItem {
    model.data = data
    return &model
}

func (model *InventoryStockItem) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}