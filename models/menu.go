package models

import (
    "encoding/json"
    "errors"
)

// Model
type Menu struct {
    // 
    CreatedAt string `json:"created_at"`
    // 
    CreatedBy string `json:"created_by"`
    // 
    DeletedAt string `json:"deleted_at"`
    // 
    Id string `json:"id"`
    // 
    Items interface{} `json:"items"`
    // 
    Label string `json:"label"`
    // 
    MenuKey string `json:"menu_key"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model Menu) New(data []byte) *Menu {
    model.data = data
    return &model
}

func (model *Menu) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}