package models

import (
    "encoding/json"
    "errors"
)

// ProductGridFilter model.
type ProductGridFilter struct {
    // The attribute code to filter on.
    Code string `json:"code"`
    // The attribute's i18n labels, for the filter's own caption.
    Label interface{} `json:"label"`
    // Which control the filter asks for — the same widget vocabulary the
    // columns use.
    Type string `json:"type"`

    // Used by Decode() method
    data []byte
}

func (model ProductGridFilter) New(data []byte) *ProductGridFilter {
    model.data = data
    return &model
}

func (model *ProductGridFilter) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}