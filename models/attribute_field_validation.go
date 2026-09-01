package models

import (
    "encoding/json"
    "errors"
)

// AttributeFieldValidation The limits the value has to satisfy, ready to hand
// to a form validator. Only the seven keys below are republished; anything
// else the tenant stored in `attributes.validation` stays there.
type AttributeFieldValidation struct {
    // Largest permitted number.
    Max float64 `json:"max"`
    // Most entries.
    MaxItems int `json:"max_items"`
    // Longest permitted text.
    MaxLength int `json:"max_length"`
    // Smallest permitted number, for a number or measure field.
    Min float64 `json:"min"`
    // Fewest entries, for a multi-select or a collection.
    MinItems int `json:"min_items"`
    // Shortest permitted text.
    MinLength int `json:"min_length"`
    // A regular expression the text has to match.
    Pattern string `json:"pattern"`

    // Used by Decode() method
    data []byte
}

func (model AttributeFieldValidation) New(data []byte) *AttributeFieldValidation {
    model.data = data
    return &model
}

func (model *AttributeFieldValidation) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}