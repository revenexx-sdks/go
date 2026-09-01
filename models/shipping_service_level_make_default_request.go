package models

import (
    "encoding/json"
    "errors"
)

// ShippingServiceLevelMakeDefaultRequest No payload — send {}.
type ShippingServiceLevelMakeDefaultRequest struct {

    // Used by Decode() method
    data []byte
}

func (model ShippingServiceLevelMakeDefaultRequest) New(data []byte) *ShippingServiceLevelMakeDefaultRequest {
    model.data = data
    return &model
}

func (model *ShippingServiceLevelMakeDefaultRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}