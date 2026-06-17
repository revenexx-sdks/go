package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOnlyTitleSlugStatusMetaAndBundleAreAppliedOtherKeysAreIgnored
// Model
type PageUpdateRequest struct {
    // 
    Bundle string `json:"bundle"`
    // 
    Meta interface{} `json:"meta"`
    // 
    Slug string `json:"slug"`
    // 
    Status string `json:"status"`
    // 
    Title string `json:"title"`

    // Used by Decode() method
    data []byte
}

func (model PageUpdateRequest) New(data []byte) *PageUpdateRequest {
    model.data = data
    return &model
}

func (model *PageUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}