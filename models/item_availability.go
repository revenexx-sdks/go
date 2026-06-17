package models

import (
    "encoding/json"
    "errors"
)

// Model
type ItemAvailability struct {
    // 
    Available float64 `json:"available"`
    // 
    Locations []interface{} `json:"locations"`
    // 
    OnHand float64 `json:"on_hand"`
    // 
    Orderable bool `json:"orderable"`
    // 
    ProductId string `json:"product_id"`
    // 
    Requested float64 `json:"requested"`
    // 
    Reserved float64 `json:"reserved"`
    // 
    Sku string `json:"sku"`
    // false = unknown to inventory; the storefront decides whether untracked
    // items sell freely.
    Tracked bool `json:"tracked"`

    // Used by Decode() method
    data []byte
}

func (model ItemAvailability) New(data []byte) *ItemAvailability {
    model.data = data
    return &model
}

func (model *ItemAvailability) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}