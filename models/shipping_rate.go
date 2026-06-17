package models

import (
    "encoding/json"
    "errors"
)

// Model
type ShippingRate struct {
    // 
    Carrier string `json:"carrier"`
    // 
    Code string `json:"code"`
    // 
    Currency string `json:"currency"`
    // 
    Description string `json:"description"`
    // 
    EtaDaysMax int `json:"eta_days_max"`
    // 
    EtaDaysMin int `json:"eta_days_min"`
    // 
    FreeReason string `json:"free_reason"`
    // 
    Labels interface{} `json:"labels"`
    // 
    Name string `json:"name"`
    // 
    Position int `json:"position"`
    // 
    Price float64 `json:"price"`
    // 
    PricingType string `json:"pricing_type"`
    // Shipping method tax class (or market default).
    TaxClass string `json:"tax_class"`
    // Tax rate % from markets.tax_classes for this market + tax_class.
    TaxRate float64 `json:"tax_rate"`

    // Used by Decode() method
    data []byte
}

func (model ShippingRate) New(data []byte) *ShippingRate {
    model.data = data
    return &model
}

func (model *ShippingRate) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}