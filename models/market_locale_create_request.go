package models

import (
    "encoding/json"
    "errors"
)

// TheOwningMarketComesFromTheRoutePathMarketId Model
type MarketLocaleCreateRequest struct {
    // Locale code, e.g. 'de-DE' (unique per market).
    Code string `json:"code"`
    // ISO 3166-1 alpha-2 country code.
    Country string `json:"country"`
    // 
    IsDefault bool `json:"is_default"`
    // ISO 639-1 language code.
    Language string `json:"language"`
    // Sort position (default 0).
    Position int `json:"position"`

    // Used by Decode() method
    data []byte
}

func (model MarketLocaleCreateRequest) New(data []byte) *MarketLocaleCreateRequest {
    model.data = data
    return &model
}

func (model *MarketLocaleCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}