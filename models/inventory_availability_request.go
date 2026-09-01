package models

import (
    "encoding/json"
    "errors"
)

// InventoryAvailabilityRequest model.
type InventoryAvailabilityRequest struct {
    // The items to check, at most 200 in one call. A cart, a category page, a
    // feed row — one call answers them all, which is why this route is the
    // batch one.
    Items []InventoryAvailabilityItem `json:"items"`
    // Restrict the check to ONE location, by its code — the stock a
    // click-and-collect store can promise today. Omitted, every ENABLED location
    // is summed; a disabled one is never counted either way.
    LocationCode string `json:"location_code"`
    // Inline single-item form: the product to move, instead of a one-entry
    // `items` array. The two forms are equivalent — nothing downstream knows
    // which arrived.
    ProductId string `json:"product_id"`
    // Inline single-item form: how many are wanted (default 1). It decides
    // `orderable` and nothing else.
    Quantity float64 `json:"quantity"`
    // Inline single-item form: the article number to move (instead of
    // `product_id`).
    Sku string `json:"sku"`

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