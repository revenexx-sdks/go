package models

import (
    "encoding/json"
    "errors"
)

// Model
type Reservation struct {
    // 
    CreatedAt string `json:"created_at"`
    // 
    ExpiresAt string `json:"expires_at"`
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
    Sku string `json:"sku"`
    // 
    Status string `json:"status"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model Reservation) New(data []byte) *Reservation {
    model.data = data
    return &model
}

func (model *Reservation) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}