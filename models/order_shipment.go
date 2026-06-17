package models

import (
    "encoding/json"
    "errors"
)

// Model
type OrderShipment struct {
    // 
    Carrier string `json:"carrier"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    Metadata interface{} `json:"metadata"`
    // 
    Number string `json:"number"`
    // 
    OrderId string `json:"order_id"`
    // 
    ShippedAt string `json:"shipped_at"`
    // 
    TrackingCode string `json:"tracking_code"`
    // 
    TrackingUrl string `json:"tracking_url"`

    // Used by Decode() method
    data []byte
}

func (model OrderShipment) New(data []byte) *OrderShipment {
    model.data = data
    return &model
}

func (model *OrderShipment) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}