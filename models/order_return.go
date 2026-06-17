package models

import (
    "encoding/json"
    "errors"
)

// Model
type OrderReturn struct {
    // 
    CompletedAt string `json:"completed_at"`
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
    Positions interface{} `json:"positions"`
    // 
    Reason string `json:"reason"`
    // 
    ReceivedAt string `json:"received_at"`
    // 
    RegisteredAt string `json:"registered_at"`
    // 
    RejectedAt string `json:"rejected_at"`
    // 
    Resolution string `json:"resolution"`
    // 
    Status string `json:"status"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model OrderReturn) New(data []byte) *OrderReturn {
    model.data = data
    return &model
}

func (model *OrderReturn) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}