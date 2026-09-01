package models

import (
    "encoding/json"
    "errors"
)

// PriceListMakeDefaultRequest No payload — send {}.
type PriceListMakeDefaultRequest struct {

    // Used by Decode() method
    data []byte
}

func (model PriceListMakeDefaultRequest) New(data []byte) *PriceListMakeDefaultRequest {
    model.data = data
    return &model
}

func (model *PriceListMakeDefaultRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}