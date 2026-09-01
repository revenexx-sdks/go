package models

import (
    "encoding/json"
    "errors"
)

// PageTemplateUpdateRequest Partial update — omitted fields keep their
// current value. A template is a COPY source, so changing it never reaches
// the pages already made from it.
type PageTemplateUpdateRequest struct {
    // A sentence about when to reach for it, shown next to the label.
    Description string `json:"description"`
    // The field this template is offered in. Null offers it in every field.
    FieldName string `json:"field_name"`
    // Whether a new page of this bundle starts from this template.
    IsDefault bool `json:"is_default"`
    // What the template is called in the picker.
    Label string `json:"label"`
    // The page type this template is offered on. Null offers it on every page
    // type.
    PageBundle string `json:"page_bundle"`
    // The blocks the template inserts, in order. Replaces the stored tree
    // completely.
    Tree []PageBlockTree `json:"tree"`

    // Used by Decode() method
    data []byte
}

func (model PageTemplateUpdateRequest) New(data []byte) *PageTemplateUpdateRequest {
    model.data = data
    return &model
}

func (model *PageTemplateUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}