package models

import (
    "encoding/json"
    "errors"
)

// MarketLocaleKeys The read and write keys for one of the market's locales,
// already resolved from the two settings.
type MarketLocaleKeys struct {
    // The market's locale this entry is about.
    Code string `json:"code"`
    // Its language part, which is also the key under language granularity.
    Language string `json:"language"`
    // Keys to try in order until one holds text. Always starts at the exact code:
    // a fallback fills a gap, it never outranks a stored value.
    Read []string `json:"read"`
    // A key inside a labels bag: a full locale ('de-DE') under regional
    // granularity, a bare language ('de') under language granularity.
    Write string `json:"write"`

    // Used by Decode() method
    data []byte
}

func (model MarketLocaleKeys) New(data []byte) *MarketLocaleKeys {
    model.data = data
    return &model
}

func (model *MarketLocaleKeys) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}