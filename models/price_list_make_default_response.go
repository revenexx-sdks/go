package models

import (
    "encoding/json"
    "errors"
)

// PriceListMakeDefaultResponse The list as it now stands, plus whoever lost
// the flag.
type PriceListMakeDefaultResponse struct {
    // Codes of the lists that lost the flag — empty when this list already held
    // it, which is what makes a repeated call free.
    Demoted []string `json:"demoted"`
    // A price list: one currency, one tax basis, one validity window, one buyer
    // scope — and the entries that price items in it. Which list wins for a
    // given buyer is decided by scope first, then priority, then the default
    // flag; see prices.resolve.
    PriceList PriceList `json:"price_list"`

    // Used by Decode() method
    data []byte
}

func (model PriceListMakeDefaultResponse) New(data []byte) *PriceListMakeDefaultResponse {
    model.data = data
    return &model
}

func (model *PriceListMakeDefaultResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}