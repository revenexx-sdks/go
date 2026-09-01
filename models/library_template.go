package models

import (
    "encoding/json"
    "errors"
)

// LibraryTemplate model.
type LibraryTemplate struct {
    // 
    BodyHtml string `json:"body_html"`
    // 
    BodyText string `json:"body_text"`
    // 
    Channel string `json:"channel"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Description string `json:"description"`
    // 
    Design []interface{} `json:"design"`
    // 
    Id string `json:"id"`
    // 
    Key string `json:"key"`
    // 
    Locale string `json:"locale"`
    // 
    Subject string `json:"subject"`
    // 
    SuggestedEvent string `json:"suggested_event"`
    // 
    SuggestedRecipient string `json:"suggested_recipient"`
    // 
    Title string `json:"title"`
    // 
    UpdatedAt string `json:"updated_at"`
    // 
    Variables []interface{} `json:"variables"`

    // Used by Decode() method
    data []byte
}

func (model LibraryTemplate) New(data []byte) *LibraryTemplate {
    model.data = data
    return &model
}

func (model *LibraryTemplate) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}