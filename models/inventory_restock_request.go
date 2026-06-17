package models

import (
    "encoding/json"
    "errors"
)

// Model
type InventoryRestockRequest struct {
    // The returned items (at most 200).
    Items []InventoryStockItem `json:"items"`
    // Restocking location (default 'main').
    LocationCode string `json:"location_code"`
    // Originating order (ledger reference).
    OrderRef string `json:"order_ref"`
    // Ledger note (e.g. return reason).
    Reason string `json:"reason"`

    // Used by Decode() method
    data []byte
}

func (model InventoryRestockRequest) New(data []byte) *InventoryRestockRequest {
    model.data = data
    return &model
}

func (model *InventoryRestockRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}