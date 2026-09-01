package models

import (
    "encoding/json"
    "errors"
)

// AttributeSchemaGroup model.
type AttributeSchemaGroup struct {
    // The group code, which is what every field in the section carries as its
    // `group`.
    Code string `json:"code"`
    // The section heading, resolved for the requested locale.
    Label string `json:"label"`
    // Where the section sits, ascending. The array is already in this order.
    Position int `json:"position"`

    // Used by Decode() method
    data []byte
}

func (model AttributeSchemaGroup) New(data []byte) *AttributeSchemaGroup {
    model.data = data
    return &model
}

func (model *AttributeSchemaGroup) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}