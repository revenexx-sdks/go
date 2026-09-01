package models

import (
    "encoding/json"
    "errors"
)

// MarketCloneCopied Child rows copied from the source, per collection. A flag
// left false is a zero here, and so is a source that had none of that kind.
type MarketCloneCopied struct {
    // Traded currencies copied from the source market.
    Currencies int `json:"currencies"`
    // Locales copied from the source market.
    Locales int `json:"locales"`
    // Tax classes copied from the source market.
    TaxClasses int `json:"tax_classes"`

    // Used by Decode() method
    data []byte
}

func (model MarketCloneCopied) New(data []byte) *MarketCloneCopied {
    model.data = data
    return &model
}

func (model *MarketCloneCopied) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}