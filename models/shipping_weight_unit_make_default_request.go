package models

import (
    "encoding/json"
    "errors"
)

// ShippingWeightUnitMakeDefaultRequest No payload — send {}.
type ShippingWeightUnitMakeDefaultRequest struct {

    // Used by Decode() method
    data []byte
}

func (model ShippingWeightUnitMakeDefaultRequest) New(data []byte) *ShippingWeightUnitMakeDefaultRequest {
    model.data = data
    return &model
}

func (model *ShippingWeightUnitMakeDefaultRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}