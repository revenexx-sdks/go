package models

import (
    "encoding/json"
    "errors"
)

// AnItemNeedsAnIdentityNameOrSku Model
type CartItemCreateRequest struct {
    // Free-form configuration — configured lines never merge.
    Configuration interface{} `json:"configuration"`
    // Defaults to the cart's currency.
    Currency string `json:"currency"`
    // Free-form metadata.
    Metadata interface{} `json:"metadata"`
    // Falls back to 'sku' when omitted.
    Name string `json:"name"`
    // 
    Position int `json:"position"`
    // 
    ProductId string `json:"product_id"`
    // Default 1.
    Quantity float64 `json:"quantity"`
    // 
    Sku string `json:"sku"`
    // Loose product snapshot at add-time (price, name, image, …).
    Snapshot interface{} `json:"snapshot"`
    // 
    TaxRate float64 `json:"tax_rate"`
    // Line type (default 'product'). Plain product lines merge by product+price;
    // configurations always stand alone.
    Type string `json:"type"`
    // 
    Unit string `json:"unit"`
    // Per-unit net price — line_total is always derived.
    UnitPrice float64 `json:"unit_price"`

    // Used by Decode() method
    data []byte
}

func (model CartItemCreateRequest) New(data []byte) *CartItemCreateRequest {
    model.data = data
    return &model
}

func (model *CartItemCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}