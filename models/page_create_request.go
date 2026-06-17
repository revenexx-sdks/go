package models

import (
    "encoding/json"
    "errors"
)

// Model
type PageCreateRequest struct {
    // 
    Bundle string `json:"bundle"`
    // 
    HostOptions interface{} `json:"hostOptions"`
    // 
    Meta interface{} `json:"meta"`
    // 
    Slug string `json:"slug"`
    // 
    SourceLanguage string `json:"sourceLanguage"`
    // 
    Title string `json:"title"`

    // Used by Decode() method
    data []byte
}

func (model PageCreateRequest) New(data []byte) *PageCreateRequest {
    model.data = data
    return &model
}

func (model *PageCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}