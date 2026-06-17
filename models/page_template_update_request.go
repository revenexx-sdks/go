package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type PageTemplateUpdateRequest struct {
    // 
    Description string `json:"description"`
    // 
    FieldName string `json:"field_name"`
    // 
    IsDefault bool `json:"is_default"`
    // 
    Label string `json:"label"`
    // 
    PageBundle string `json:"page_bundle"`
    // Serialized block trees ({ bundle, props, props_i18n, options, children }).
    Tree []interface{} `json:"tree"`

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