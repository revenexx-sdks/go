package models

import (
    "encoding/json"
    "errors"
)

// AttributeSchemaFamily The family the fields belong to, or null when none
// was named — then the answer is every attribute of the `entity_type`,
// which is what a reference entity or an asset family has instead of a
// family.
type AttributeSchemaFamily struct {
    // The family's code — the value `?family_code=` takes.
    Code string `json:"code"`
    // The family's id.
    Id string `json:"id"`
    // The family name, resolved for the requested locale.
    Label string `json:"label"`
    // Which of these fields is the product's display name.
    LabelAttribute string `json:"label_attribute"`

    // Used by Decode() method
    data []byte
}

func (model AttributeSchemaFamily) New(data []byte) *AttributeSchemaFamily {
    model.data = data
    return &model
}

func (model *AttributeSchemaFamily) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}