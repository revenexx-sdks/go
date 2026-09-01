package models

import (
    "encoding/json"
    "errors"
)

// PriceTier One rung of the winning list’s quantity ladder for this item.
type PriceTier struct {
    // The quantity this rung applies from. The rung with the highest
    // `quantity_min` at or below the requested quantity is the one `unit_price`
    // on the item was taken from.
    QuantityMin float64 `json:"quantity_min"`
    // Unit of measure the rung’s price is per. Absent when the entry names
    // none.
    Unit string `json:"unit"`
    // The rung’s price for ONE unit, in the answer’s `currency` and on the
    // item’s `tax_basis` — decimal major units, exactly as stored. Tiers are
    // NOT tax-adjusted: only the chosen price gets
    // `unit_price_net`/`unit_price_gross`.
    UnitPrice float64 `json:"unit_price"`

    // Used by Decode() method
    data []byte
}

func (model PriceTier) New(data []byte) *PriceTier {
    model.data = data
    return &model
}

func (model *PriceTier) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}