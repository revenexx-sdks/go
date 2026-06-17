package models

import (
    "encoding/json"
    "errors"
)

// AnItemToCheckProductIdOrSku Model
type InventoryAvailabilityItem struct {
    // 
    ProductId string `json:"product_id"`
    // Requested quantity for the orderable check (default 1).
    Quantity float64 `json:"quantity"`
    // 
    Sku string `json:"sku"`

    // Used by Decode() method
    data []byte
}

func (model InventoryAvailabilityItem) New(data []byte) *InventoryAvailabilityItem {
    model.data = data
    return &model
}

func (model *InventoryAvailabilityItem) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}