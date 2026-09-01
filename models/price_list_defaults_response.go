package models

import (
    "encoding/json"
    "errors"
)

// PriceListDefaultsResponse What seeding found and what it had to write.
// Idempotent twice over: by code, and by the existence of ANY default list
// — so changing default_price_list_code later never produces a second
// default.
type PriceListDefaultsResponse struct {
    // Codes of the lists this call created — empty on a tenant that was already
    // seeded.
    Created []string `json:"created"`
    // Codes of the lists that were already there, so nothing was written for
    // them.
    Existing []string `json:"existing"`

    // Used by Decode() method
    data []byte
}

func (model PriceListDefaultsResponse) New(data []byte) *PriceListDefaultsResponse {
    model.data = data
    return &model
}

func (model *PriceListDefaultsResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}