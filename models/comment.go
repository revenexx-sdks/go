package models

import (
    "encoding/json"
    "errors"
)

// Model
type Comment struct {
    // 
    AuthorId string `json:"author_id"`
    // 
    AuthorName string `json:"author_name"`
    // 
    BlockUuids interface{} `json:"block_uuids"`
    // 
    Body string `json:"body"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    PageId string `json:"page_id"`
    // 
    ParentId string `json:"parent_id"`
    // 
    Resolved bool `json:"resolved"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model Comment) New(data []byte) *Comment {
    model.data = data
    return &model
}

func (model *Comment) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}