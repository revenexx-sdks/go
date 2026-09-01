package models

import (
    "encoding/json"
    "errors"
)

// ProductsBatchRequest Name the products either way, or both ways. Send at
// least one non-empty list; the two are unioned and a product named twice
// comes back once.
type ProductsBatchRequest struct {
    // Product ids, when the caller already holds them.
    Ids []string `json:"ids"`
    // Product SKUs — the identifier a foreign system carries, which is why this
    // route exists at all.
    Skus []string `json:"skus"`

    // Used by Decode() method
    data []byte
}

func (model ProductsBatchRequest) New(data []byte) *ProductsBatchRequest {
    model.data = data
    return &model
}

func (model *ProductsBatchRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}