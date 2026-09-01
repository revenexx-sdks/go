package models

import (
    "encoding/json"
    "errors"
)

// PriceEntriesLadderResponse The generated ladder as stored, plus the
// rounding policy that shaped it.
type PriceEntriesLadderResponse struct {
    // The generated rungs, one per requested quantity, ascending — this IS the
    // item's ladder in this list.
    Entries []PriceEntry `json:"entries"`
    // Decimals each tier was rounded to before snapping — the tenant's
    // price_precision.
    Precision int `json:"precision"`
    // true when the item's existing entries in this list were removed first (the
    // default), so the answer is the whole ladder rather than an addition to one.
    Replaced bool `json:"replaced"`
    // The price ending each tier was snapped to — the request's, or the
    // tenant's bulk_adjust_rounding.
    Rounding string `json:"rounding"`
    // How they landed on the last decimal — the tenant's rounding_mode.
    RoundingMode string `json:"rounding_mode"`

    // Used by Decode() method
    data []byte
}

func (model PriceEntriesLadderResponse) New(data []byte) *PriceEntriesLadderResponse {
    model.data = data
    return &model
}

func (model *PriceEntriesLadderResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}