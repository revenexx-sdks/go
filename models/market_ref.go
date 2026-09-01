package models

import (
    "encoding/json"
    "errors"
)

// MarketRef The market that was read from, resolved — so a caller who
// passed a code back gets the uuid, and one who passed a uuid gets the code
// the rest of the platform stores.
type MarketRef struct {
    // The source market's code — the value other apps scope by.
    Code string `json:"code"`
    // The source market's primary key.
    Id string `json:"id"`

    // Used by Decode() method
    data []byte
}

func (model MarketRef) New(data []byte) *MarketRef {
    model.data = data
    return &model
}

func (model *MarketRef) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}