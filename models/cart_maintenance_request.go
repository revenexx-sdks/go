package models

import (
    "encoding/json"
    "errors"
)

// CartMaintenanceRequest model.
type CartMaintenanceRequest struct {
    // Report what the sweep WOULD do and write nothing. Worth doing before a
    // first retention run: cart_ttl_days deletes carts and their lines.
    DryRun bool `json:"dry_run"`

    // Used by Decode() method
    data []byte
}

func (model CartMaintenanceRequest) New(data []byte) *CartMaintenanceRequest {
    model.data = data
    return &model
}

func (model *CartMaintenanceRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}