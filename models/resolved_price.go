package models

import (
    "encoding/json"
    "errors"
)

// Model
type ResolvedPrice struct {
    // 
    Currency string `json:"currency"`
    // 
    LineTotal float64 `json:"line_total"`
    // true = no price for this buyer context — show "price on request", never
    // 0.
    OnRequest bool `json:"on_request"`
    // 
    PriceList interface{} `json:"price_list"`
    // 
    ProductId string `json:"product_id"`
    // 
    Quantity float64 `json:"quantity"`
    // 
    Sku string `json:"sku"`
    // Resolved tax class code (from the product, or the market default).
    TaxClass string `json:"tax_class"`
    // 
    TaxIncluded bool `json:"tax_included"`
    // Tax rate % from markets.tax_classes for this market + tax_class.
    TaxRate float64 `json:"tax_rate"`
    // 
    Tiers []interface{} `json:"tiers"`
    // Stored price as-is (net or gross per tax_included). Prefer
    // unit_price_net/unit_price_gross.
    UnitPrice float64 `json:"unit_price"`
    // Gross unit price (incl. tax).
    UnitPriceGross float64 `json:"unit_price_gross"`
    // Net unit price (excl. tax).
    UnitPriceNet float64 `json:"unit_price_net"`

    // Used by Decode() method
    data []byte
}

func (model ResolvedPrice) New(data []byte) *ResolvedPrice {
    model.data = data
    return &model
}

func (model *ResolvedPrice) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}