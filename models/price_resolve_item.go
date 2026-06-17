package models

import (
    "encoding/json"
    "errors"
)

// IdentifyByProductIdOrSkuAnItemWithoutIdentityResolvesToOnRequestWithAPerItemError
// Model
type PriceResolveItem struct {
    // Product to price.
    ProductId string `json:"product_id"`
    // Requested quantity for tier selection and line_total (default 1;
    // non-positive values fall back to 1).
    Quantity float64 `json:"quantity"`
    // SKU to price (alternative to product_id).
    Sku string `json:"sku"`

    // Used by Decode() method
    data []byte
}

func (model PriceResolveItem) New(data []byte) *PriceResolveItem {
    model.data = data
    return &model
}

func (model *PriceResolveItem) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}