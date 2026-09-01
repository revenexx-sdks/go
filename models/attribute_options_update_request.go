package models

import (
    "encoding/json"
    "errors"
)

// AttributeOptionsUpdateRequest Partial update — omitted fields keep their
// current value.
type AttributeOptionsUpdateRequest struct {
    // The select / multi-select attribute these are the permitted values of.
    // Deleting the attribute deletes its options with it.
    AttributeId string `json:"attribute_id"`
    // The value actually STORED in a record's `attribute_values` when this option
    // is picked — never the label. Unique within the attribute.
    Code string `json:"code"`
    // What the option is called, per language tag. Two tenants may label the same
    // code differently; only the code is ever written into a record.
    Labels interface{} `json:"labels"`
    // Order in the dropdown, ascending. Options that tie keep the order the
    // database returns them in, so give every option a position if the order
    // matters.
    Position int `json:"position"`
    // A colour or texture chip for the picker. Null for an option that is not
    // visual.
    Swatch interface{} `json:"swatch"`

    // Used by Decode() method
    data []byte
}

func (model AttributeOptionsUpdateRequest) New(data []byte) *AttributeOptionsUpdateRequest {
    model.data = data
    return &model
}

func (model *AttributeOptionsUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}