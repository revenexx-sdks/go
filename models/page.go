package models

import (
    "encoding/json"
    "errors"
)

// Model
type Page struct {
    // 
    AnalyzeIgnored interface{} `json:"analyze_ignored"`
    // 
    Bundle string `json:"bundle"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    CreatedBy string `json:"created_by"`
    // 
    DeletedAt string `json:"deleted_at"`
    // 
    HostOptions interface{} `json:"host_options"`
    // 
    Id string `json:"id"`
    // 
    Meta interface{} `json:"meta"`
    // 
    PublishedRevisionId string `json:"published_revision_id"`
    // 
    Slug string `json:"slug"`
    // 
    SourceLanguage string `json:"source_language"`
    // 
    Status string `json:"status"`
    // 
    Title string `json:"title"`
    // 
    UpdatedAt string `json:"updated_at"`
    // 
    UpdatedBy string `json:"updated_by"`

    // Used by Decode() method
    data []byte
}

func (model Page) New(data []byte) *Page {
    model.data = data
    return &model
}

func (model *Page) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}