package models

import (
    "encoding/json"
    "errors"
)

// InventoryShipTo Where the order is going. Read ONLY when the tenant's
// `allocation_strategy` is 'nearest' — under 'priority' or
// 'single_location' it is accepted and ignored, so sending it is never wrong,
// it is just not always heard.
type InventoryShipTo struct {
    // ISO country code of the delivery address. Locations whose `address.country`
    // matches are tried before the rest, which is what stops a German order
    // pulling from an overseas warehouse that merely sorts first.
    Country string `json:"country"`
    // Prefer this location above everything else — a click-and-collect store
    // the customer picked. It is a preference, not a demand: if it cannot cover
    // the item the allocator moves on to the next location.
    LocationCode string `json:"location_code"`

    // Used by Decode() method
    data []byte
}

func (model InventoryShipTo) New(data []byte) *InventoryShipTo {
    model.data = data
    return &model
}

func (model *InventoryShipTo) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}