package models

import (
    "encoding/json"
    "errors"
)

// ShippingTaxContext Tax resolution status of this answer. resolved=false ⇒
// tax_class/tax_rate are unknown, NOT zero.
type ShippingTaxContext struct {
    // The market whose tax classes were applied.
    MarketId string `json:"market_id"`
    // Human-readable form of `reason`, safe to log or show an operator. One
    // sentence per reason; the example is the `no_markets` wording.
    Message string `json:"message"`
    // Only when resolved=false — why no rate could be applied.
    Reason string `json:"reason"`
    // Whether a tax rate could be applied at all. FALSE means every rate's
    // tax_class and tax_rate are UNKNOWN — not zero, and not tax-free. A
    // checkout that adds 0 % on this is wrong; read `reason` and either ask for a
    // market or refuse to quote.
    Resolved bool `json:"resolved"`
    // Where the market came from: 'request' (market_id), 'header'
    // (x-revenexx-market), 'country' (the market matching the destination) or
    // 'sole_market' (the tenant has exactly one).
    Source string `json:"source"`
    // Present when the market is known but registers no tax classes and the
    // tenant's default_shipping_tax_rate supplied the number instead.
    Via string `json:"via"`

    // Used by Decode() method
    data []byte
}

func (model ShippingTaxContext) New(data []byte) *ShippingTaxContext {
    model.data = data
    return &model
}

func (model *ShippingTaxContext) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}