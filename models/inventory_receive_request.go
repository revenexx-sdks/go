package models

import (
    "encoding/json"
    "errors"
)

// Model
type InventoryReceiveRequest struct {
    // The inbound items (at most 200).
    Items []InventoryStockItem `json:"items"`
    // Receiving location (default 'main').
    LocationCode string `json:"location_code"`
    // Ledger note (e.g. delivery note number).
    Reason string `json:"reason"`

    // Used by Decode() method
    data []byte
}

func (model InventoryReceiveRequest) New(data []byte) *InventoryReceiveRequest {
    model.data = data
    return &model
}

func (model *InventoryReceiveRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}