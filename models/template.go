package models

import (
    "encoding/json"
    "errors"
)

// Model
type Template struct {
    // 
    CreatedAt string `json:"created_at"`
    // 
    CreatedBy string `json:"created_by"`
    // 
    Description string `json:"description"`
    // 
    FieldName string `json:"field_name"`
    // 
    Id string `json:"id"`
    // 
    IsDefault bool `json:"is_default"`
    // 
    Label string `json:"label"`
    // 
    PageBundle string `json:"page_bundle"`
    // 
    Tree interface{} `json:"tree"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model Template) New(data []byte) *Template {
    model.data = data
    return &model
}

func (model *Template) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}