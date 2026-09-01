package models

import (
    "encoding/json"
    "errors"
)

// AttributeFieldOption model.
type AttributeFieldOption struct {
    // What to show in the picker, already resolved for the requested locale.
    Label string `json:"label"`
    // Colour/texture chip, when the option carries one — `{"hex": "#c0c0c0"}`.
    Swatch interface{} `json:"swatch"`
    // The stored value — an `attribute_options.code`, or a
    // `reference_entity_records.code` when the options ARE a reference entity.
    // This, never the label, is what goes into `attribute_values`.
    Value string `json:"value"`

    // Used by Decode() method
    data []byte
}

func (model AttributeFieldOption) New(data []byte) *AttributeFieldOption {
    model.data = data
    return &model
}

func (model *AttributeFieldOption) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}