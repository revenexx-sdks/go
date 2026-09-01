package models

import (
    "encoding/json"
    "errors"
)

// DeliveryPage One published page resolved for one language, ready to render:
// i18n fallback applied per field, blocks outside their publish window
// removed, library references expanded inline.
type DeliveryPage struct {
    // The page's block tree, keyed by field name — `{ "content": [ … ] }`. A
    // theme renders the field it knows and ignores the rest.
    Fields interface{} `json:"fields"`
    // The page frame — everything a theme needs before it starts rendering
    // blocks.
    Page interface{} `json:"page"`

    // Used by Decode() method
    data []byte
}

func (model DeliveryPage) New(data []byte) *DeliveryPage {
    model.data = data
    return &model
}

func (model *DeliveryPage) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}