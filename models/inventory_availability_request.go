package models

import (
    "encoding/json"
    "errors"
)

// Model
type InventoryAvailabilityRequest struct {
    // The items to check (batch, at most 200).
    Items []InventoryAvailabilityItem `json:"items"`
    // Restrict the check to one location (default: all enabled locations).
    LocationCode string `json:"location_code"`

    // Used by Decode() method
    data []byte
}

func (model InventoryAvailabilityRequest) New(data []byte) *InventoryAvailabilityRequest {
    model.data = data
    return &model
}

func (model *InventoryAvailabilityRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}