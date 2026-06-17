package models

import (
    "encoding/json"
    "errors"
)

// TemplateVariable Model
type TemplateVariable struct {
    // Variable Description.
    Description string `json:"description"`
    // Variable Name.
    Name string `json:"name"`
    // Variable Placeholder.
    Placeholder string `json:"placeholder"`
    // Is the variable required?
    Required bool `json:"required"`
    // Variable secret flag. Secret variables can only be updated or deleted, but
    // never read.
    Secret bool `json:"secret"`
    // Variable Type.
    Type string `json:"type"`
    // Variable Value.
    Value string `json:"value"`

    // Used by Decode() method
    data []byte
}

func (model TemplateVariable) New(data []byte) *TemplateVariable {
    model.data = data
    return &model
}

func (model *TemplateVariable) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}