package models

import (
    "encoding/json"
    "errors"
)

// PriceTaxContext Tax resolution status of this answer. resolved=false ⇒
// tax_class/tax_rate are unknown, NOT zero.
type PriceTaxContext struct {
    // The market whose tax classes were applied.
    MarketId string `json:"market_id"`
    // Human-readable form of `reason`, in English. Safe to log; not phrased for a
    // buyer.
    Message string `json:"message"`
    // Only when resolved=false — why no rate could be applied.
    Reason string `json:"reason"`
    // true ⇒ every priced item carries `tax_class`, `tax_rate`,
    // `unit_price_net` and `unit_price_gross`. false ⇒ those are null because
    // the rate could not be established — read `reason`, and never as "no tax
    // due".
    Resolved bool `json:"resolved"`
    // Where the market came from: 'request' (market_id), 'header'
    // (x-revenexx-market) or 'sole_market' (the tenant has exactly one).
    Source string `json:"source"`

    // Used by Decode() method
    data []byte
}

func (model PriceTaxContext) New(data []byte) *PriceTaxContext {
    model.data = data
    return &model
}

func (model *PriceTaxContext) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}