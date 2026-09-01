package models

import (
    "encoding/json"
    "errors"
)

// PageTemplateCreateRequest The blocks to freeze, and where the template
// should be offered.
type PageTemplateCreateRequest struct {
    // A sentence about when to reach for it.
    Description string `json:"description"`
    // The field this template should be offered in. Null offers it in every
    // field.
    FieldName string `json:"fieldName"`
    // Whether a new page of that type should start from this template.
    IsDefault bool `json:"isDefault"`
    // What the template is called in the picker.
    Label string `json:"label"`
    // The page type this template should be offered on. Omit to take the current
    // page's own type.
    PageBundle string `json:"pageBundle"`
    // The blocks to serialize into the template, each with its whole subtree.
    // They are read from the CURRENT edit state, so unpublished changes are
    // included.
    Uuids []string `json:"uuids"`

    // Used by Decode() method
    data []byte
}

func (model PageTemplateCreateRequest) New(data []byte) *PageTemplateCreateRequest {
    model.data = data
    return &model
}

func (model *PageTemplateCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}