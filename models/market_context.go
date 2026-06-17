package models

import (
    "encoding/json"
    "errors"
)

// Model
type MarketContext struct {
    // 
    Currencies []MarketCurrency `json:"currencies"`
    // 
    Locales []MarketLocale `json:"locales"`
    // 
    Market Market `json:"market"`
    // 
    TaxClasses []MarketTaxClass `json:"tax_classes"`

    // Used by Decode() method
    data []byte
}

func (model MarketContext) New(data []byte) *MarketContext {
    model.data = data
    return &model
}

func (model *MarketContext) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}