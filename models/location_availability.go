package models

import (
    "encoding/json"
    "errors"
)

// LocationAvailability What one location holds of this item. Only enabled
// locations appear, and only those with a stock row for the item — a
// location that has never held it is absent rather than zero.
type LocationAvailability struct {
    // on_hand − reserved at this location — what this one place can still
    // promise.
    Available float64 `json:"available"`
    // The location CODE (`locations.code`) — the same value `location_code`
    // takes in a request. Falls back to the raw location id in the rare case
    // where the location row disappeared between the two reads.
    Location string `json:"location"`
    // Physically at this location, promised units included.
    OnHand float64 `json:"on_hand"`
    // Held for orders at this location.
    Reserved float64 `json:"reserved"`

    // Used by Decode() method
    data []byte
}

func (model LocationAvailability) New(data []byte) *LocationAvailability {
    model.data = data
    return &model
}

func (model *LocationAvailability) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}