package models

import (
    "encoding/json"
    "errors"
)

// APositionOfThePlacedOrderNeedsAnIdentityNameOrSkuItemsAreSNAPSHOTSCarryTheProductCopyPricesAreFrozenAtPlaceTime
// Model
type OrderItemCreateRequest struct {
    // Free-form configuration of configured lines.
    Configuration interface{} `json:"configuration"`
    // 
    CostCenter string `json:"cost_center"`
    // Free-form metadata.
    Metadata interface{} `json:"metadata"`
    // Falls back to 'sku' when omitted.
    Name string `json:"name"`
    // Explicit position number; otherwise numbered in steps of the order range's
    // position_step.
    Position int `json:"position"`
    // 
    PositionText string `json:"position_text"`
    // Frozen product snapshot at place-time ('snapshot' is accepted as an alias).
    Product interface{} `json:"product"`
    // 
    ProductId string `json:"product_id"`
    // Default 1.
    Quantity float64 `json:"quantity"`
    // 
    Sku string `json:"sku"`
    // Alias for 'product'.
    Snapshot interface{} `json:"snapshot"`
    // Derived from line_total and tax_rate when omitted.
    TaxAmount float64 `json:"tax_amount"`
    // Percent (default 0).
    TaxRate float64 `json:"tax_rate"`
    // Line type (default 'product').
    Type string `json:"type"`
    // 
    Unit string `json:"unit"`
    // Per-unit net price — line_total is always derived.
    UnitPrice float64 `json:"unit_price"`
    // Free-form user data.
    UserData interface{} `json:"user_data"`

    // Used by Decode() method
    data []byte
}

func (model OrderItemCreateRequest) New(data []byte) *OrderItemCreateRequest {
    model.data = data
    return &model
}

func (model *OrderItemCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}