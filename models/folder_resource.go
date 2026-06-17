package models

import (
    "encoding/json"
    "errors"
)

// Model
type FolderResource struct {
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    IsSystem bool `json:"is_system"`
    // 
    Name string `json:"name"`
    // 
    ParentId string `json:"parent_id"`
    // 
    Path string `json:"path"`
    // 
    TenantId string `json:"tenant_id"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model FolderResource) New(data []byte) *FolderResource {
    model.data = data
    return &model
}

func (model *FolderResource) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}