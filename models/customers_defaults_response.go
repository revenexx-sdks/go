package models

import (
    "encoding/json"
    "errors"
)

// CustomersDefaultsResponse model.
type CustomersDefaultsResponse struct {
    // One entry per value set, keyed by its route name — `payment-terms`,
    // `address-types`, `lifecycle-stages`, `contact-event-kinds`. Each says what
    // THIS call did: `created` are the codes it inserted, `existing` the seeded
    // codes it found already there and left completely alone (a merchant's rename
    // included). A second call therefore answers with everything under `existing`
    // and nothing under `created`.
    Sets interface{} `json:"sets"`

    // Used by Decode() method
    data []byte
}

func (model CustomersDefaultsResponse) New(data []byte) *CustomersDefaultsResponse {
    model.data = data
    return &model
}

func (model *CustomersDefaultsResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}