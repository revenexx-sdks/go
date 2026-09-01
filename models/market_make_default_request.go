package models

import (
    "encoding/json"
    "errors"
)

// MarketMakeDefaultRequest No payload — send {}. Which market is promoted
// comes from the path, and there is nothing else to say.
type MarketMakeDefaultRequest struct {

    // Used by Decode() method
    data []byte
}

func (model MarketMakeDefaultRequest) New(data []byte) *MarketMakeDefaultRequest {
    model.data = data
    return &model
}

func (model *MarketMakeDefaultRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}