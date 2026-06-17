package models

import (
    "encoding/json"
    "errors"
)

// AnItemAndItsSIGNEDCorrectionProductIdOrSku Model
type InventoryAdjustItem struct {
    // 
    ProductId string `json:"product_id"`
    // Signed delta (±on_hand) — must be non-zero.
    Quantity float64 `json:"quantity"`
    // 
    Sku string `json:"sku"`

    // Used by Decode() method
    data []byte
}

func (model InventoryAdjustItem) New(data []byte) *InventoryAdjustItem {
    model.data = data
    return &model
}

func (model *InventoryAdjustItem) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}