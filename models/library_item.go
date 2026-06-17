package models

import (
    "encoding/json"
    "errors"
)

// Model
type LibraryItem struct {
    // 
    Bundle string `json:"bundle"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    CreatedBy string `json:"created_by"`
    // 
    DeletedAt string `json:"deleted_at"`
    // 
    Id string `json:"id"`
    // 
    Label string `json:"label"`
    // 
    Tree interface{} `json:"tree"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model LibraryItem) New(data []byte) *LibraryItem {
    model.data = data
    return &model
}

func (model *LibraryItem) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}