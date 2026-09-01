package models

import (
    "encoding/json"
    "errors"
)

// ShippingTaxClassUsage What in this app still points at a market tax class,
// by code.
type ShippingTaxClassUsage struct {
    // The tax-class code that was asked about, echoed back.
    Code string `json:"code"`
    // True when this market's shipping_tax_class setting names the code — the
    // class every method that names none falls back to.
    FallbackSetting bool `json:"fallback_setting"`
    // True when at least one method or the market fallback setting names it. The
    // single field a caller deciding whether to allow a delete needs; the rest is
    // so it can word the refusal.
    InUse bool `json:"in_use"`
    // The first 20 of them, so a refusal can name names instead of a number.
    Methods []interface{} `json:"methods"`
    // How many methods name this code as their own tax_class. Capped at 500 — a
    // tenant with more shipping methods than that has a bigger problem than an
    // imprecise count.
    ShippingMethods int `json:"shipping_methods"`

    // Used by Decode() method
    data []byte
}

func (model ShippingTaxClassUsage) New(data []byte) *ShippingTaxClassUsage {
    model.data = data
    return &model
}

func (model *ShippingTaxClassUsage) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}